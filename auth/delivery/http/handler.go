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
		data := &entity.ConfirmationUserRequest{}

		err := json.NewDecoder(request.Body).Decode(&data)
		if err != nil {
			json.NewEncoder(writer).Encode("error")
		}
		confirmError := ah.useCase.ConfirmUser(data)

		if confirmError != nil {
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(entity.NewErrorResponse(confirmError))
			return
		}

		json.NewEncoder(writer).Encode(entity.NewSuccessResponse(struct{}{}))
	}
}
