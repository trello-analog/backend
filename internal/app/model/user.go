package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	TwoAuth  bool   `json:"two_auth"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
	)
}
