package auth

import (
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
)

type AuthRepository interface {
	CreateUser(user *model.User) (*entity.IdResponse, *customerrors.APIError)
	GetUserById(id int) (*model.User, *customerrors.APIError)
	GetUserByField(field string, value interface{}) (*model.User, *customerrors.APIError)
	CreateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError
	GetConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError)
	DeleteConfirmationCode(id int) *customerrors.APIError
	UpdateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError
	GetConfirmationCodesByField(field string, value interface{}) ([]model.ConfirmationCode, *customerrors.APIError)
	GetLastConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError)
}
