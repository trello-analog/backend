package http

import (
	"encoding/json"
	"github.com/trello-analog/backend/auth"
	"github.com/trello-analog/backend/entity"
	model "github.com/trello-analog/backend/models"
	"net/http"
)

type AuthHandler struct {
	useCase auth.UseCase
}

func NewAuthHandler(useCase auth.UseCase) *AuthHandler {
	return &AuthHandler{
		useCase: useCase,
	}
}

func (ah *AuthHandler) SignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user := &model.User{}
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			json.NewEncoder(writer).Encode("error")
		}

		result, signUpError := ah.useCase.SignUp(user)

		if signUpError != nil {
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(entity.NewErrorResponse(signUpError))
			return
		}

		json.NewEncoder(writer).Encode(entity.NewSuccessResponse(result))
	}
}

func (ah *AuthHandler) ConfirmUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		code := request.FormValue("code")
		email := request.FormValue("email")

		err := ah.useCase.ConfirmUser(&entity.ConfirmationUserRequest{
			Code:  code,
			Email: email,
		})

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(entity.NewErrorResponse(err))
			return
		}

		json.NewEncoder(writer).Encode(entity.NewSuccessResponse(struct{}{}))
	}
}
