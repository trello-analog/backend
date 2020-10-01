package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/services"
)

type UserRole struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	TwoAuth   bool   `json:"twoAuth"`
	Avatar    string `json:"avatar"`
	TokenCode string `json:"token_code"`
}

type UserForFrontend struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Login   string `json:"login"`
	TwoAuth bool   `json:"twoAuth"`
	Avatar  string `json:"avatar"`
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

func (u *User) SetPassword(password string) *User {
	u.Password = password
	return u
}

func (u *User) CryptPassword() *User {
	ps := services.PasswordService{}
	u.Password = ps.GetCryptPassword(u.Password)
	return u
}

func (u *User) IsPasswordEqual(password string) bool {
	ps := services.PasswordService{}
	return ps.ComparePasswords(password, u.Password)
}

func (u *User) IsEmpty() bool {
	return u.ID == 0
}

func (u *User) IsTwoAuth() bool {
	return u.TwoAuth
}

func (u *User) CreateTokenCode() *User {
	code, _ := uuid.NewV4()
	u.TokenCode = code.String()
	return u
}

func (u *User) ToClientStruct() *UserForFrontend {
	return &UserForFrontend{
		ID:      u.ID,
		Email:   u.Email,
		Login:   u.Login,
		TwoAuth: u.TwoAuth,
		Avatar:  u.Avatar,
	}
}
