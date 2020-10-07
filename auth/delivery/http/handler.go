package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trello-analog/backend/auth"
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
	"github.com/trello-analog/backend/services"
	"net/http"
)

type AuthHandler struct {
	useCase  auth.UseCase
	response *services.ResponseService
}

func NewAuthHandler(useCase auth.UseCase) *AuthHandler {
	return &AuthHandler{
		useCase:  useCase,
		response: services.NewResponseService(),
	}
}

func (ah *AuthHandler) SignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user := &model.User{}
		err := json.NewDecoder(request.Body).Decode(&user)

		if err != nil {
			ah.response.SetWriter(writer).SetData(customerrors.ParseError).Error()
			return
		}

		result, signUpError := ah.useCase.SignUp(user)

		if signUpError != nil {
			ah.response.SetWriter(writer).SetData(signUpError).Error()
			return
		}
		ah.response.SetWriter(writer).SetData(result).Success()
	}
}

func (ah *AuthHandler) ConfirmUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &entity.ConfirmationUserRequest{}
		err := json.NewDecoder(request.Body).Decode(&data)

		if err != nil {
			ah.response.SetWriter(writer).SetData(customerrors.ParseError).Error()
			return
		}

		confirm, confirmError := ah.useCase.ConfirmUser(data)

		if confirmError != nil {
			ah.response.SetWriter(writer).SetData(confirmError).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(confirm).Success()
	}
}

func (ah *AuthHandler) ResendConfirmationCode() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &auth.ResendConfirmationCodeRequest{}
		err := json.NewDecoder(request.Body).Decode(&data)

		if err != nil {
			ah.response.SetWriter(writer).SetData(customerrors.ParseError).Error()
			return
		}

		sendError := ah.useCase.ResendConfirmationCode(data.Email)

		if sendError != nil {
			ah.response.SetWriter(writer).SetData(sendError).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(struct{}{}).Success()
	}
}

func (ah *AuthHandler) ForgotPassword() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &auth.ForgotPasswordRequest{}
		err := json.NewDecoder(request.Body).Decode(&data)

		if err != nil {
			ah.response.SetWriter(writer).SetData(customerrors.ParseError).Error()
			return
		}

		forgotError := ah.useCase.ForgotPassword(data.Email)

		if forgotError != nil {
			ah.response.SetWriter(writer).SetData(forgotError).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(struct{}{}).Success()
	}
}

func (ah *AuthHandler) CheckForgotPasswordCode() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := mux.Vars(request)["id"]
		check, err := ah.useCase.CheckForgotPassword(id)

		if err != nil {
			ah.response.SetWriter(writer).SetData(err).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(check).Success()
	}
}

func (ah *AuthHandler) RestorePassword() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &auth.RestorePasswordRequest{}
		err := json.NewDecoder(request.Body).Decode(&data)

		if err != nil {
			return
		}

		restore, restoreErr := ah.useCase.RestorePassword(data)

		if restoreErr != nil {
			ah.response.SetWriter(writer).SetData(restoreErr).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(restore).Success()
	}
}

func (ah *AuthHandler) SignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &auth.SignInRequest{}
		err := json.NewDecoder(request.Body).Decode(&data)

		if err != nil {
			return
		}

		token, twoAuth, signInErr := ah.useCase.SignIn(data)

		if signInErr != nil {
			ah.response.SetWriter(writer).SetData(signInErr).Error()
			return
		}

		if token == nil && twoAuth != nil {
			ah.response.SetWriter(writer).SetData(twoAuth).Success()
			return
		}

		ah.response.SetWriter(writer).SetData(token).Success()
	}
}

func (ah *AuthHandler) ResendTwoAuth() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		type requestStruct struct {
			UserId int `json:"userId"`
		}
		data := &requestStruct{}
		err := json.NewDecoder(request.Body).Decode(&data)

		if err != nil {
			return
		}

		_, code, resendError := ah.useCase.ResendTwoAuthCode(data.UserId)

		if resendError != nil {
			ah.response.SetWriter(writer).SetData(resendError).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(code).Success()
	}
}

func (ah *AuthHandler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token := &entity.Token{
			AccessToken:  request.Header.Get("access-token"),
			RefreshToken: request.Header.Get("refresh-token"),
		}

		user, err := ah.useCase.Login(token)

		if err != nil {
			ah.response.SetWriter(writer).SetData(err).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(user).Success()
	}
}

func (ah *AuthHandler) Logout() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token := &entity.Token{
			AccessToken:  request.Header.Get("access-token"),
			RefreshToken: request.Header.Get("refresh-token"),
		}

		err := ah.useCase.Logout(token)

		if err != nil {
			ah.response.SetWriter(writer).SetData(err).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(struct{}{}).Success()
	}
}

func (ah *AuthHandler) SendTwoAuth() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := &auth.TwoAuthCodeRequest{}
		err := json.NewDecoder(request.Body).Decode(&data)

		if err != nil {
			return
		}

		twoAuth, twoAuthError := ah.useCase.SendTwoAuth(data)

		if twoAuthError != nil {
			ah.response.SetWriter(writer).SetData(twoAuthError).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(twoAuth).Success()
	}
}

func (ah *AuthHandler) RefreshToken() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token := &entity.Token{
			AccessToken:  request.Header.Get("access-token"),
			RefreshToken: request.Header.Get("refresh-token"),
		}

		tokenResp, err := ah.useCase.RefreshToken(token)

		if err != nil {
			ah.response.SetWriter(writer).SetData(err).Error()
			return
		}

		ah.response.SetWriter(writer).SetData(tokenResp).Success()
	}
}
