package services

import (
	"github.com/trello-analog/backend/config"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func (ps *PasswordService) GetCryptPassword(rawPassword string) string {
	secret := config.GetConfig().PasswordSecret
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword+secret), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(cryptedPassword)
}

func (ps *PasswordService) ComparePasswords(rawPassword, cryptedPassword string) bool {
	secret := config.GetConfig().PasswordSecret
	err := bcrypt.CompareHashAndPassword([]byte(cryptedPassword), []byte(rawPassword+secret))
	if err != nil {
		return false
	}
	return true
}
