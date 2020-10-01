package emails

type SignUpEmail struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Host  string `json:"host"`
	Code  string `json:"code"`
}

type TwoAuthEmail struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Email string `json:"email"`
}
