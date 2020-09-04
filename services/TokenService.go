package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/trello-analog/backend/config"
	"github.com/trello-analog/backend/entity"
	"time"
)

type TokenData struct {
	UserId int
}

type TokenClaims struct {
	jwt.StandardClaims
	Data *TokenData `json:"data"`
}

type TokenService struct {
	Token *entity.Token
}

func (t *TokenService) GenerateToken(data *TokenData) error {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 30)),
		},
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24 * 7)),
		},
	})

	accessTokenString, accessError := accessToken.SignedString([]byte(config.GetConfig().AccessTokenSecret))
	refreshTokenString, refreshError := refreshToken.SignedString([]byte(config.GetConfig().RefreshTokenSecret))

	if accessError != nil {
		return errors.New("Access token error")
	}

	if refreshError != nil {
		return errors.New("Refresh token error")
	}

	t.Token = &entity.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return nil
}

func (t *TokenService) GetToken() *entity.Token {
	return t.Token
}

func (t *TokenService) ParseToken(tokenString, mode string) (*TokenData, error) {
	claims := &TokenClaims{}
	if mode == "access" {
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().AccessTokenSecret), nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().RefreshTokenSecret), nil
		})
		if err != nil {
			return nil, err
		}
	}
	return claims.Data, nil
}
