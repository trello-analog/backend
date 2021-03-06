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
	RequiredForFilling            = NewAPIError(http.StatusBadRequest, 1, "Обязательно для заполнения!")
	UserLoginAlreadyExists        = NewAPIError(http.StatusBadRequest, 2, "Пользователь с таким логином уже существует!")
	UserEmailAlreadyExists        = NewAPIError(http.StatusBadRequest, 3, "Пользователь с таким email-ом уже существует!")
	InvalidEmail                  = NewAPIError(http.StatusBadRequest, 4, "Невалидный email!")
	InvalidPassword               = NewAPIError(http.StatusBadRequest, 5, "Невалидный пароль!")
	NotFound                      = NewAPIError(http.StatusNotFound, 6, "Не найдено!")
	PostServerError               = NewAPIError(http.StatusBadRequest, 7, "Ошибка почтового сервера!")
	CodeExpired                   = NewAPIError(http.StatusBadRequest, 8, "Время жизни этого кода подтверждения истёк!")
	ParseError                    = NewAPIError(http.StatusBadRequest, 9, "Ошибка при запросе")
	EmailNotExists                = NewAPIError(http.StatusNotFound, 10, "Данный e-mail на зарегистрирован!")
	ForgotPasswordCodeIsNotActive = NewAPIError(http.StatusBadRequest, 11, "Даннная ссылка больше не активна!")
	EnteredCurrentPassword        = NewAPIError(http.StatusBadRequest, 12, "Вы ввели текущий пароль!")
	PasswordsAreNotEqual          = NewAPIError(http.StatusBadRequest, 13, "Пароли не совпадают!")
	WrongCredential               = NewAPIError(http.StatusBadRequest, 14, "Неверные данные для входа!")
	UserNotFound                  = NewAPIError(http.StatusNotFound, 15, "Юзер не найден!")
	UserNotConfirmed              = NewAPIError(http.StatusNotFound, 16, "Юзер не подтвержден!")
	CodeNotFound                  = NewAPIError(http.StatusNotFound, 17, "Код не найден!")
	TwoAuthCodeAlreadyExists      = NewAPIError(http.StatusNotFound, 18, "Код уже существует!")
	WrongTwoAuthCode              = NewAPIError(http.StatusNotFound, 19, "Неверный код!")
	TokenExpired                  = NewAPIError(http.StatusUnauthorized, 20, "Время жизни токена истекло!")
)
