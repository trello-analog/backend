package usecase

import (
	"github.com/trello-analog/backend/auth"
	"github.com/trello-analog/backend/config"
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/emails"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
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
	userByLogin, _ := auc.repository.GetUserByField("login", &user.Login)
	userByEmail, _ := auc.repository.GetUserByField("email", &user.Email)
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

	user.CryptPassword()
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

	if forgotPasswordCode.IsConfirmed() || forgotPasswordCode.IsCodeExpired() || forgotPasswordCode.IsRelevant() {
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
