package auth

import (
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
)

type UseCase interface {
	SignUp(user *model.User) (*entity.IdResponse, *customerrors.APIError)
	ConfirmUser(data *entity.ConfirmationUserRequest) (*ConfirmUserResponse, *customerrors.APIError)
	ResendConfirmationCode(email string) *customerrors.APIError
	ForgotPassword(email string) *customerrors.APIError
	CheckForgotPassword(code string) (*ConfirmUserResponse, *customerrors.APIError)
	RestorePassword(data *RestorePasswordRequest) (*ConfirmUserResponse, *customerrors.APIError)
	SignIn(data *SignInRequest) (*SignInResponseToken, *SignInResponseTwoAuth, *customerrors.APIError)
	ResendTwoAuthCode(userId int) (*SignInResponseToken, *SignInResponseTwoAuth, *customerrors.APIError)
	SendTwoAuth(data *TwoAuthCodeRequest) (*SignInResponseToken, *customerrors.APIError)
	Login(token *entity.Token) (*model.UserForFrontend, *customerrors.APIError)
	Logout(token *entity.Token) *customerrors.APIError
	RefreshToken(token *entity.Token) (*SignInResponseToken, *customerrors.APIError)
}
