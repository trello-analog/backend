package entity

type ConfirmationUserRequest struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}
