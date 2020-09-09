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

func (a *AuthRepository) userTable() *gorm.DB {
	return a.db.Table("users")
}

func (a *AuthRepository) confirmationCodeTable() *gorm.DB {
	return a.db.Table("confirmation_codes")
}

func (a *AuthRepository) forgotPasswordTable() *gorm.DB {
	return a.db.Table("forgot_password")
}

func (a *AuthRepository) CreateUser(user *model.User) (*entity.IdResponse, *customerrors.APIError) {
	u := a.userTable().Create(&user)

	if u.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusBadRequest, 10, u.Error.Error())
	}

	return &entity.IdResponse{
		ID: user.ID,
	}, nil
}

func (a *AuthRepository) GetUserById(id int) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.userTable().First(&user, id)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return user, nil
}

func (a *AuthRepository) GetUserByField(field string, value interface{}) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.userTable().Where(field+" = ?", value).Find(&user)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return user, nil
}

func (a *AuthRepository) UpdateUser(user *model.User) *customerrors.APIError {
	result := a.userTable().Model(&user).Updates(map[string]interface{}{
		"email":    user.Email,
		"login":    user.Login,
		"password": user.Password,
		"avatar":   user.Avatar,
		"two_auth": user.TwoAuth,
	})

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) CreateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError {
	result := a.confirmationCodeTable().Create(&code)

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) GetConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError) {
	code := &model.ConfirmationCode{}
	result := a.confirmationCodeTable().Where(field+" = ?", value).Find(&code)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return code, nil
}

func (a *AuthRepository) DeleteConfirmationCode(id int) *customerrors.APIError {
	code := &model.ConfirmationCode{}
	result := a.confirmationCodeTable().Delete(&code, id)

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) UpdateConfirmationCode(code *model.ConfirmationCode) *customerrors.APIError {
	result := a.confirmationCodeTable().Model(&code).Updates(map[string]interface{}{
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
	result := a.confirmationCodeTable().Where(field+" = ?", value).Find(&codes)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return codes, nil
}

func (a *AuthRepository) GetLastConfirmationCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError) {
	code := &model.ConfirmationCode{}
	result := a.confirmationCodeTable().Where(field+" = ?", value).Last(&code)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return code, nil
}

func (a *AuthRepository) CreateForgotPasswordCode(code *model.ConfirmationCode) *customerrors.APIError {
	result := a.forgotPasswordTable().Create(&code)

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) GetForgotPasswordCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError) {
	code := &model.ConfirmationCode{}
	result := a.forgotPasswordTable().Where(field+" = ?", value).Find(&code)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return code, nil
}

func (a *AuthRepository) GetLastForgotPasswordCodeByField(field string, value interface{}) (*model.ConfirmationCode, *customerrors.APIError) {
	code := &model.ConfirmationCode{}
	result := a.forgotPasswordTable().Where(field+" = ?", value).Last(&code)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return code, nil
}

func (a *AuthRepository) UpdateForgotPasswordCode(code *model.ConfirmationCode) *customerrors.APIError {
	result := a.forgotPasswordTable().Model(&code).Updates(map[string]interface{}{
		"code":      code.Code,
		"user_id":   code.UserId,
		"expired":   code.Expired,
		"confirmed": code.Confirmed,
		"last":      code.Last,
	})

	if result.Error != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return nil
}

func (a *AuthRepository) CountForgotPasswordCodes(field string, value interface{}) (int64, *customerrors.APIError) {
	codes := []model.ConfirmationCode{}
	result := a.forgotPasswordTable().Where(field+" = ?", value).Find(&codes)

	if result.Error != nil {
		return 0, customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return result.RowsAffected, nil
}
