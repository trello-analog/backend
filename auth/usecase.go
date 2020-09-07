package auth

import (
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
)

type UseCase interface {
	SignUp(user *model.User) (*entity.IdResponse, *customerrors.APIError)
	ConfirmUser(data *entity.ConfirmationUserRequest) *customerrors.APIError
}
