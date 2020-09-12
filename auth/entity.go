package auth

import "github.com/trello-analog/backend/entity"

type SignInRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInResponseToken struct {
	Token *entity.Token `json:"token"`
}

type SignInResponseTwoAuth struct {
	CheckCode string `json:"check_code"`
	Expired   string `json:"expired"`
}

type RestorePasswordRequest struct {
	Code              string `json:"code"`
	NewPassword       string `json:"newPassword"`
	RepeatNewPassword string `json:"repeatNewPassword"`
}

type TwoAuthCodeRequest struct {
	CheckCode string `json:"check_code"`
	Code      string `json:"code"`
}

type ConfirmUserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResendConfirmationCodeRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}
