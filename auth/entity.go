package auth

type SignInRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RestorePasswordRequest struct {
	Code                  string
	NewPassword           string
	RepeatPasswordRequest string
}

type TwoAuthCodeRequest struct {
	CheckCode string
	Code      string
}

type TwoAuthCodeResponse struct {
	CheckCode string
	Expired   string
}
