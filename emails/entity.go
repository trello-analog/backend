package emails

type SignUpEmail struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Host  string `json:"host"`
	Code  string `json:"code"`
}
