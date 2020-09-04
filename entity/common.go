package entity

type IdResponse struct {
	ID int `json:"id"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
