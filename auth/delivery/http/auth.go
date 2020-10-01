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
	router.HandleFunc("/sign-in", handler.SignIn()).Methods(http.MethodPost)
	router.HandleFunc("/resend-two-auth", handler.ResendTwoAuth()).Methods(http.MethodPost)
	router.HandleFunc("/two-auth", handler.SendTwoAuth()).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.Login()).Methods(http.MethodPost)
	router.HandleFunc("/logout", handler.Logout()).Methods(http.MethodPost)
}
