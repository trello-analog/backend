package entity

type IdResponse struct {
	ID int `json:"id"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
