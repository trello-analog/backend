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
		writer.Header().Set("Content-Type", "application/json")

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
