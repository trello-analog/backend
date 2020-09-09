package auth

type SignInRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RestorePasswordRequest struct {
	Code                  string `json:"code"`
	NewPassword           string `json:"new_password"`
	RepeatPasswordRequest string `json:"repeat_password_request"`
}

type TwoAuthCodeRequest struct {
	CheckCode string `json:"check_code"`
	Code      string `json:"code"`
}

type TwoAuthCodeResponse struct {
	CheckCode string `json:"check_code"`
	Expired   string `json:"expired"`
}

type ConfirmUserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResendConfirmationCodeRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}
