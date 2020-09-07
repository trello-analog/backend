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
}

func NewAuthUseCase(repository auth.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		repository: repository,
	}
}

func (auc *AuthUseCase) SignUp(user *model.User) (*entity.IdResponse, *customerrors.APIError) {
	userByLogin, _ := auc.repository.GetUserByField("login", &user.Login)
	userByEmail, _ := auc.repository.GetUserByField("email", &user.Email)
	sender := emails.NewEmailSender()
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

	sendError := sender.SendEmailAfterSignUp(&emails.SignUpEmail{
		Email: user.Email,
		Name:  user.Login,
		Host:  cfg.Client.Host,
		Code:  confirmationCode.Code,
	})
	if sendError != nil {
		return nil, customerrors.SignUpEmailError
	}

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (auc *AuthUseCase) ConfirmUser(data *entity.ConfirmationUserRequest) *customerrors.APIError {
	code, err := auc.repository.GetConfirmationCodeByField("code", data.Code)

	if code.ID == 0 {
		return customerrors.NotFound
	}

	if code.IsCodeExpired() {
		return customerrors.CodeExpired
	}

	if err != nil {
		return err
	}

	deleteError := auc.repository.DeleteConfirmationCode(code.ID)

	if deleteError != nil {
		return deleteError
	}

	return nil
}
