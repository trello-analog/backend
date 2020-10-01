package postgres

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type AuthRepository struct {
	db *entity.Database
}

func NewAuthRepository(db *entity.Database) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) userTable() *gorm.DB {
	return a.db.Postgres.Table("users")
}

func (a *AuthRepository) confirmationCodeTable() *gorm.DB {
	return a.db.Postgres.Table("confirmation_codes")
}

func (a *AuthRepository) forgotPasswordTable() *gorm.DB {
	return a.db.Postgres.Table("forgot_password")
}

func (a *AuthRepository) GetUserByQuery(query interface{}, args ...interface{}) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.userTable().Where(query, args...).Find(&user)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	return user, nil
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
	var count int64 = 0
	result := a.forgotPasswordTable().Where(field+" = ?", value).Count(&count)

	if result.Error != nil {
		return 0, customerrors.NewAPIError(http.StatusBadRequest, 10, result.Error.Error())
	}

	return count, nil
}

func (a *AuthRepository) CreateTwoAuthCode(code string, userID int) *customerrors.APIError {
	err := a.db.Redis.Set(context.Background(), strconv.Itoa(userID), code, time.Minute*3).Err()
	if err != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, err.Error())
	}

	return nil
}

func (a *AuthRepository) GetTwoAuthCode(key string) (string, *customerrors.APIError) {
	val, err := a.db.Redis.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return "", customerrors.CodeNotFound
	}

	if err != nil {
		return "", customerrors.NewAPIError(http.StatusBadRequest, 10, err.Error())
	}

	return val, nil
}

func (a *AuthRepository) GetExpirationTwoAuthCodeByKey(key string) (float64, *customerrors.APIError) {
	exp, err := a.db.Redis.TTL(context.Background(), key).Result()
	if err != nil {
		return 0, customerrors.NewAPIError(http.StatusBadRequest, 10, err.Error())
	}
	if exp < 0 {
		return 0, customerrors.CodeNotFound
	}
	return exp.Seconds(), nil
}

func (a *AuthRepository) DeleteTwoAuthCode(key string) *customerrors.APIError {
	_, err := a.db.Redis.Del(context.Background(), key).Result()

	if err == redis.Nil {
		return customerrors.WrongTwoAuthCode
	}

	if err != nil {
		return customerrors.NewAPIError(http.StatusBadRequest, 10, err.Error())
	}

	return nil
}
