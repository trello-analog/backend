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

	if user == (&model.User{}) {
		return nil, customerrors.UserNotFound
	}

	return user, nil
}

func (a *AuthRepository) GetUserByField(field string, value interface{}) (*model.User, *customerrors.APIError) {
	user := &model.User{}
	result := a.db.Table("users").Where(field+" = ?", value).Find(&user)

	if result.Error != nil {
		return nil, customerrors.NewAPIError(http.StatusNotFound, 10, result.Error.Error())
	}

	if user == (&model.User{}) {
		return nil, customerrors.UserNotFound
	}

	return user, nil
}
