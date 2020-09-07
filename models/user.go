package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/services"
)

type UserRole struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	TwoAuth  bool   `json:"two_auth"`
	Avatar   string `json:"avatar"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(
			&u.Email,
			validation.Required.Error(customerrors.RequiredForFilling.Message),
			is.Email.Error(customerrors.InvalidEmail.Message),
		),
		validation.Field(
			&u.Login,
			validation.Required.Error(customerrors.RequiredForFilling.Message),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error(customerrors.RequiredForFilling.Message),
			validation.Length(6, 100).Error(customerrors.InvalidPassword.Message),
		),
	)
}

func (u *User) CryptPassword() *User {
	ps := services.PasswordService{}
	u.Password = ps.GetCryptPassword(u.Password)

	return u
}
