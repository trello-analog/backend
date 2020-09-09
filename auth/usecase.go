package auth

import (
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
)

type UseCase interface {
	SignUp(user *model.User) (*entity.IdResponse, *customerrors.APIError)
	ConfirmUser(data *entity.ConfirmationUserRequest) (*ConfirmUserResponse, *customerrors.APIError)
	ResendConfirmationCode(email string) *customerrors.APIError
	ForgotPassword(email string) *customerrors.APIError
	CheckForgotPassword(code string) (*ConfirmUserResponse, *customerrors.APIError)
}
