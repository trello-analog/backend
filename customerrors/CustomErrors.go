package customerrors

import (
	"net/http"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAPIError(status, code int, message string) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
		Code:    code,
	}
}

var (
	RequiredForFilling     = NewAPIError(http.StatusBadRequest, 1, "Обязательно для заполнения!")
	UserLoginAlreadyExists = NewAPIError(http.StatusBadRequest, 2, "Пользователь с таким логином уже существует!")
	UserEmailAlreadyExists = NewAPIError(http.StatusBadRequest, 3, "Пользователь с таким email-ом уже существует!")
	InvalidEmail           = NewAPIError(http.StatusBadRequest, 4, "Невалидный email!")
	InvalidPassword        = NewAPIError(http.StatusBadRequest, 5, "Невалидный пароль!")
	NotFound               = NewAPIError(http.StatusNotFound, 6, "Не найдено!")
	PostServerError        = NewAPIError(http.StatusBadRequest, 7, "Ошибка почтового сервера!")
	CodeExpired            = NewAPIError(http.StatusBadRequest, 8, "Время жизни этого кода подтверждения истёк!")
	ParseError             = NewAPIError(http.StatusBadRequest, 9, "Ошибка при запросе")
	EmailNotExists         = NewAPIError(http.StatusNotFound, 10, "Данный e-mail на зарегистрирован!")
)
