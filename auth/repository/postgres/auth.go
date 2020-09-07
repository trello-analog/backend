package postgres

import (
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
	"gorm.io/gorm"
	"net/http"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) CreateUser(user *model.User) (*entity.IdResponse, *customerrors.APIError) {
	u := a.db.Table("users").Create(&user)

	if u.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusBadRequest, 10, u.Error.Error())
	}

	return &entity.IdResponse{
		ID: user.ID,
	}, nil
}

func (a *AuthRepository) GetUserById(id int) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.db.Table("users").First(&user, id)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return user, nil
}

func (a *AuthRepository) GetUserByField(field string, value interface{}) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.db.Table("users").Where(field+" = ?", value).Find(&user)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return user, nil
}

func (a *AuthRepository) CreateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError {
	result := a.db.Table("confirmation_codes").Create(&code)

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) GetConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError) {
	code := &model.ConfirmationCode{}
	result := a.db.Table("confirmation_codes").Where(field+" = ?", value).Find(&code)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return code, nil
}

func (a *AuthRepository) DeleteConfirmationCode(id int) *customerrors.APIError {
	code := &model.ConfirmationCode{}
	result := a.db.Table("users").First(&code, id)

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return nil
}
