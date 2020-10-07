package usecase

import (
	"github.com/trello-analog/backend/auth"
	"github.com/trello-analog/backend/config"
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/emails"
	"github.com/trello-analog/backend/entity"
	"github.com/trello-analog/backend/helper"
	model "github.com/trello-analog/backend/models"
	"github.com/trello-analog/backend/services"
	"net/http"
)

type AuthUseCase struct {
	repository auth.AuthRepository
	sender     *emails.EmailSender
}

func NewAuthUseCase(repository auth.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		repository: repository,
		sender:     emails.NewEmailSender(),
	}
}

func (auc *AuthUseCase) SignUp(user *model.User) (*entity.IdResponse, *customerrors.APIError) {
	userByLogin, _ := auc.repository.GetUserByQuery("login = ?", user.Login)
	userByEmail, _ := auc.repository.GetUserByQuery("email = ?", user.Email)
	cfg := config.GetConfig()

	if userByLogin.ID != 0 {
		return nil, customerrors.UserLoginAlreadyExists
	}
	if userByEmail.ID != 0 {
		return nil, customerrors.UserEmailAlreadyExists
	}
	if user.Validate() != nil {
		return nil, customerrors.NewAPIError(http.StatusBadRequest, 8, user.Validate().Error())
	}

	user.CryptPassword().CreateTokenCode()
	result, err := auc.repository.CreateUser(user)
	confirmationCode := model.NewConfirmationCode(result.ID)
	codeErr := auc.repository.CreateConfirmationCode(confirmationCode)

	if codeErr != nil {
		return nil, codeErr
	}

	sendError := auc.sender.SendEmailAfterSignUp(&emails.SignUpEmail{
		Email: user.Email,
		Name:  user.Login,
		Host:  cfg.Client.Host,
		Code:  confirmationCode.Code,
	})

	if sendError != nil {
		return nil, customerrors.PostServerError
	}

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (auc *AuthUseCase) ConfirmUser(data *entity.ConfirmationUserRequest) (*auth.ConfirmUserResponse, *customerrors.APIError) {
	code, err := auc.repository.GetConfirmationCodeByField("code", data.Code)
	if err != nil {
		return nil, err
	}

	if code.ID == 0 {
		return nil, customerrors.NotFound
	}

	if code.IsCodeExpired() {
		return &auth.ConfirmUserResponse{
			Status:  "error",
			Message: customerrors.CodeExpired.Message,
		}, nil
	}

	if code.IsConfirmed() {
		return &auth.ConfirmUserResponse{
			Status:  "info",
			Message: "Аккаунт уже был подтврждён!",
		}, nil
	}

	if !code.IsLast() {
		return &auth.ConfirmUserResponse{
			Status:  "error",
			Message: "Этот код подтверждения уже неактивен. Возможно, Вы создавали новые коды",
		}, nil
	}

	code.MakeConfirmed()
	updateError := auc.repository.UpdateConfirmationCode(code)

	if updateError != nil {
		return nil, updateError
	}

	return &auth.ConfirmUserResponse{
		Status:  "success",
		Message: "Аккаунт верифицирован!",
	}, nil
}

func (auc *AuthUseCase) ResendConfirmationCode(email string) *customerrors.APIError {
	user, err := auc.repository.GetUserByField("email", email)
	cfg := config.GetConfig()

	if err != nil {
		return err
	}

	currentCode, currentCodeErr := auc.repository.GetLastConfirmationCodeByField("user_id", user.ID)

	if currentCodeErr != nil {
		return currentCodeErr
	}

	currentCode.MakeIrrelevant()
	updateCurrentError := auc.repository.UpdateConfirmationCode(currentCode)

	if updateCurrentError != nil {
		return updateCurrentError
	}

	newCode := model.NewConfirmationCode(user.ID)
	newCodeErr := auc.repository.CreateConfirmationCode(newCode)

	if newCodeErr != nil {
		return newCodeErr
	}

	sendError := auc.sender.SendEmailAfterSignUp(&emails.SignUpEmail{
		Email: user.Email,
		Name:  user.Login,
		Host:  cfg.Client.Host,
		Code:  newCode.Code,
	})
	if sendError != nil {
		return customerrors.PostServerError
	}

	return nil
}

func (auc *AuthUseCase) ForgotPassword(email string) *customerrors.APIError {
	user, err := auc.repository.GetUserByField("email", email)
	forgotPasswordCodesCount, _ := auc.repository.CountForgotPasswordCodes("user_id", user.ID)

	if forgotPasswordCodesCount > 0 {
		currentCode, _ := auc.repository.GetLastForgotPasswordCodeByField("user_id", user.ID)
		currentCode.MakeIrrelevant()
		updateErr := auc.repository.UpdateForgotPasswordCode(currentCode)

		if updateErr != nil {
			return updateErr
		}
	}

	cfg := config.GetConfig()

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return customerrors.EmailNotExists
	}

	code := model.NewConfirmationCode(user.ID)
	createError := auc.repository.CreateForgotPasswordCode(code)

	if createError != nil {
		return createError
	}

	sendError := auc.sender.SendForgotPasswordEmail(&emails.SignUpEmail{
		Email: user.Email,
		Name:  user.Login,
		Host:  cfg.Client.Host,
		Code:  code.Code,
	})

	if sendError != nil {
		return customerrors.PostServerError
	}

	return nil
}

func (auc *AuthUseCase) CheckForgotPassword(code string) (*auth.ConfirmUserResponse, *customerrors.APIError) {
	forgotPasswordCode, err := auc.repository.GetForgotPasswordCodeByField("code", code)

	if err != nil {
		return nil, err
	}

	if forgotPasswordCode.IsEmpty() {
		return nil, customerrors.NotFound
	}

	if forgotPasswordCode.IsConfirmed() || forgotPasswordCode.IsCodeExpired() || !forgotPasswordCode.IsRelevant() {
		return &auth.ConfirmUserResponse{
			Status:  "error",
			Message: "Данная ссылка уже нективна!",
		}, nil
	}

	return &auth.ConfirmUserResponse{
		Status:  "success",
		Message: "",
	}, nil
}

func (auc *AuthUseCase) RestorePassword(data *auth.RestorePasswordRequest) (*auth.ConfirmUserResponse, *customerrors.APIError) {
	forgotPasswordCode, err := auc.repository.GetForgotPasswordCodeByField("code", data.Code)

	if err != nil {
		return nil, err
	}

	if forgotPasswordCode.IsEmpty() {
		return nil, customerrors.NotFound
	}

	if forgotPasswordCode.IsConfirmed() || forgotPasswordCode.IsCodeExpired() || !forgotPasswordCode.IsRelevant() {
		return nil, customerrors.ForgotPasswordCodeIsNotActive
	}

	user, userErr := auc.repository.GetUserById(forgotPasswordCode.UserId)

	if userErr != nil {
		return nil, userErr
	}

	if user.IsPasswordEqual(data.NewPassword) {
		return nil, customerrors.EnteredCurrentPassword
	}

	if data.NewPassword != data.RepeatNewPassword {
		return nil, customerrors.PasswordsAreNotEqual
	}

	user.SetPassword(data.NewPassword).CryptPassword()
	forgotPasswordCode.MakeConfirmed()
	updateErr := auc.repository.UpdateUser(user)

	if updateErr != nil {
		return nil, updateErr
	}

	updateForgotPasswordCodeError := auc.repository.UpdateForgotPasswordCode(forgotPasswordCode)

	if updateForgotPasswordCodeError != nil {
		return nil, updateForgotPasswordCodeError
	}

	return &auth.ConfirmUserResponse{
		Status:  "success",
		Message: "",
	}, nil
}

func (auc *AuthUseCase) SignIn(data *auth.SignInRequest) (*auth.SignInResponseToken, *auth.SignInResponseTwoAuth, *customerrors.APIError) {
	user, err := auc.repository.GetUserByQuery("email = ? OR login = ?", data.Name, data.Name)

	if err != nil {
		return nil, nil, err
	}

	if user.IsEmpty() {
		return nil, nil, customerrors.UserNotFound
	}

	if !user.IsPasswordEqual(data.Password) {
		return nil, nil, customerrors.WrongCredential
	}

	code, _ := auc.repository.GetLastConfirmationCodeByField("user_id", user.ID)

	if !code.IsConfirmed() {
		return nil, nil, customerrors.UserNotConfirmed
	}

	if user.IsTwoAuth() {
		return auc.handleTwoAuthCode(user)
	}

	token := services.NewTokenService(&services.TokenData{
		UserId:   user.ID,
		TempCode: user.TokenCode,
	})

	return &auth.SignInResponseToken{
		Token: token.GetToken(),
	}, nil, nil
}

func (auc *AuthUseCase) ResendTwoAuthCode(userId int) (*auth.SignInResponseToken, *auth.SignInResponseTwoAuth, *customerrors.APIError) {
	user, userErr := auc.repository.GetUserById(userId)

	if userErr != nil {
		return nil, nil, userErr
	}
	return auc.handleTwoAuthCode(user)
}

func (auc *AuthUseCase) handleTwoAuthCode(user *model.User) (*auth.SignInResponseToken, *auth.SignInResponseTwoAuth, *customerrors.APIError) {
	code, codeErr := auc.repository.GetTwoAuthCode(helper.IntToString(user.ID))
	if codeErr != nil && codeErr != customerrors.CodeNotFound {
		return nil, nil, codeErr
	}
	if code != "" {
		exp, expErr := auc.repository.GetExpirationTwoAuthCodeByKey(helper.IntToString(user.ID))
		if expErr != nil {
			return nil, nil, expErr
		}

		return nil, &auth.SignInResponseTwoAuth{
			UserId:  user.ID,
			Expired: exp,
		}, nil
	}
	twoAuthCode := helper.GenerateTwoAuthCode()
	auc.repository.CreateTwoAuthCode(twoAuthCode, user.ID)
	sendErr := auc.sender.SendTwoAuthCode(&emails.TwoAuthEmail{
		Name:  user.Login,
		Code:  twoAuthCode,
		Email: user.Email,
	})

	if sendErr != nil {
		return nil, nil, customerrors.PostServerError
	}

	return nil, &auth.SignInResponseTwoAuth{
		UserId:  user.ID,
		Expired: 180,
	}, nil
}

func (auc *AuthUseCase) SendTwoAuth(data *auth.TwoAuthCodeRequest) (*auth.SignInResponseToken, *customerrors.APIError) {
	user, userErr := auc.repository.GetUserById(data.UserId)

	if userErr != nil {
		return nil, customerrors.UserNotFound
	}

	code, err := auc.repository.GetTwoAuthCode(helper.IntToString(data.UserId))

	if err != nil {
		return nil, err
	}

	if data.Code != code {
		return nil, customerrors.WrongTwoAuthCode
	}

	deleteCodeError := auc.repository.DeleteTwoAuthCode(helper.IntToString(data.UserId))

	if deleteCodeError != nil {
		return nil, deleteCodeError
	}

	token := services.NewTokenService(&services.TokenData{
		UserId:   user.ID,
		TempCode: user.TokenCode,
	})

	return &auth.SignInResponseToken{
		Token: token.GetToken(),
	}, nil
}

func (auc *AuthUseCase) Login(token *entity.Token) (*model.UserForFrontend, *customerrors.APIError) {
	tokenClaims, err := services.ParseToken(token.AccessToken, "access")

	if err != nil {
		return nil, err
	}

	user, userErr := auc.repository.GetUserById(tokenClaims.Data.UserId)

	if userErr != nil {
		return nil, userErr
	}

	return user.ToClientStruct(), nil
}

func (auc *AuthUseCase) Logout(token *entity.Token) *customerrors.APIError {
	tokenClaims, err := services.ParseToken(token.AccessToken, "access")

	if err != nil {
		return err
	}

	user, userErr := auc.repository.GetUserById(tokenClaims.Data.UserId)

	if userErr != nil {
		return userErr
	}
	user.CreateTokenCode()
	updateErr := auc.repository.UpdateUser(user)

	if updateErr != nil {
		return updateErr
	}

	return nil
}

func (auc *AuthUseCase) RefreshToken(token *entity.Token) (*auth.SignInResponseToken, *customerrors.APIError) {
	tokenClaims, err := services.ParseToken(token.RefreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	user, userErr := auc.repository.GetUserById(tokenClaims.Data.UserId)

	if userErr != nil {
		return nil, userErr
	}

	t := services.NewTokenService(&services.TokenData{
		UserId:   user.ID,
		TempCode: user.TokenCode,
	})

	return &auth.SignInResponseToken{
		Token: t.GetToken(),
	}, nil
}
