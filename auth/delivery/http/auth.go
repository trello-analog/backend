package http

import (
	"github.com/gorilla/mux"
	"github.com/trello-analog/backend/auth"
	"net/http"
)

func AuthEndpoints(router *mux.Router, useCase auth.UseCase) {
	handler := NewAuthHandler(useCase)

	router.HandleFunc("/sign-up", handler.SignUp()).Methods(http.MethodPost)
	router.HandleFunc("/confirm", handler.ConfirmUser()).Methods(http.MethodPost)
	router.HandleFunc("/resend-confirm", handler.ResendConfirmationCode()).Methods(http.MethodPost)
	router.HandleFunc("/forgot-password", handler.ForgotPassword()).Methods(http.MethodPost)
	router.HandleFunc("/restore-password", handler.RestorePassword()).Methods(http.MethodPost)
	router.HandleFunc("/restore-password/{id}", handler.CheckForgotPasswordCode()).Methods(http.MethodGet)
}
