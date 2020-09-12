package auth

import (
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
)

type AuthRepository interface {
	GetUserByQuery(query interface{}, args ...interface{}) (*model.User, *customerrors.APIError)
	CreateUser(user *model.User) (*entity.IdResponse, *customerrors.APIError)
	GetUserById(id int) (*model.User, *customerrors.APIError)
	GetUserByField(field string, value interface{}) (*model.User, *customerrors.APIError)
	UpdateUser(user *model.User) *customerrors.APIError

	CreateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError
	GetConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError)
	DeleteConfirmationCode(id int) *customerrors.APIError
	UpdateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError
	GetConfirmationCodesByField(field string, value interface{}) ([]model.ConfirmationCode, *customerrors.APIError)
	GetLastConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError)

	CreateForgotPasswordCode(code *model.ConfirmationCode) *customerrors.APIError
	GetForgotPasswordCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError)
	GetLastForgotPasswordCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError)
	UpdateForgotPasswordCode(code *model.ConfirmationCode) *customerrors.APIError
	CountForgotPasswordCodes(field string, value interface{}) (int64, *customerrors.APIError)

	//CreateTwoAuthCode(code *model.ConfirmationCode) *customerrors.APIError
	//GetTwoAuthCodeByQuery(query interface{}, args ...interface{}) (*model.ConfirmationCode, *customerrors.APIError)
}
