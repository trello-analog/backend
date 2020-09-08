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

func (a *AuthRepository) UserTable() *gorm.DB {
	return a.db.Table("users")
}

func (a *AuthRepository) ConfirmationCodeTable() *gorm.DB {
	return a.db.Table("confirmation_codes")
}

func (a *AuthRepository) ForgotPasswordTable() *gorm.DB {
	return a.db.Table("forgot_password")
}

func (a *AuthRepository) CreateUser(user *model.User) (*entity.IdResponse, *customerrors.APIError) {
	u := a.UserTable().Create(&user)

	if u.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusBadRequest, 10, u.Error.Error())
	}

	return &entity.IdResponse{
		ID: user.ID,
	}, nil
}

func (a *AuthRepository) GetUserById(id int) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.UserTable().First(&user, id)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return user, nil
}

func (a *AuthRepository) GetUserByField(field string, value interface{}) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.UserTable().Where(field+" = ?", value).Find(&user)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return user, nil
}

func (a *AuthRepository) CreateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError {
	result := a.ConfirmationCodeTable().Create(&code)

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) GetConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError) {
	code := &model.ConfirmationCode{}
	result := a.ConfirmationCodeTable().Where(field+" = ?", value).Find(&code)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return code, nil
}

func (a *AuthRepository) DeleteConfirmationCode(id int) *customerrors.APIError {
	code := &model.ConfirmationCode{}
	result := a.ConfirmationCodeTable().Delete(&code, id)

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) UpdateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError {
	result := a.db.Model(&code).Table("confirmation_codes").Updates(map[string]interface{}{
		"code":      code.Code,
		"user_id":   code.UserId,
		"expired":   code.Expired,
		"confirmed": code.Confirmed,
		"last":      code.Last,
	})

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) GetConfirmationCodesByField(field string, value interface{}) ([]model.ConfirmationCode, *customerrors.APIError) {
	codes := []model.ConfirmationCode{}
	result := a.ConfirmationCodeTable().Where(field+" = ?", value).Find(&codes)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return codes, nil
}

func (a *AuthRepository) GetLastConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError) {
	code := &model.ConfirmationCode{}
	result := a.ConfirmationCodeTable().Where(field+" = ?", value).Last(&code)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return code, nil
}
