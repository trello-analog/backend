package http

import (
	"github.com/gorilla/mux"
	"github.com/trello-analog/backend/auth"
	"net/http"
)

func AuthEndpoints(router *mux.Router, useCase auth.UseCase) {
	handler := NewAuthHandler(useCase)

	router.HandleFunc("/sign-up", handler.SignUp()).Methods(http.MethodPost)
	router.
		HandleFunc("/confirm", handler.SignUp()).
		Queries("code", "{code}", "email", "{email}").
		Methods(http.MethodPost)
}
