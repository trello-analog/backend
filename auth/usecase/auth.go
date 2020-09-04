package usecase

import (
	"github.com/trello-analog/backend/auth"
	"github.com/trello-analog/backend/customerrors"
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

	if err != nil {
		return nil, err
	}

	return result, nil

}
