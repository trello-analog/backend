package http

import (
	"encoding/json"
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
		}

		sendError := ah.useCase.ResendConfirmationCode(data.Email)

		if sendError != nil {
			ah.response.SetWriter(writer).SetData(sendError).Error()
		}

		ah.response.SetWriter(writer).SetData(struct{}{}).Success()
	}
}
