package model

import (
	uuid "github.com/nu7hatch/gouuid"
	"time"
)

type ConfirmationCode struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	UserId  int    `json:"user_id"`
	Expired string `json:"expired"`
}

func NewConfirmationCode(userId int) *ConfirmationCode {
	code, _ := uuid.NewV4()

	return &ConfirmationCode{
		Code:    code.String(),
		UserId:  userId,
		Expired: time.Now().Add(time.Hour * 24).UTC().String(),
	}
}

func (ccs *ConfirmationCode) IsCodeExpired() bool {
	now := time.Now().Unix()
	codeExpire, err := time.Parse(time.Now().UTC().String(), ccs.Expired)

	if err != nil {
		return false
	}

	return now > codeExpire.Unix()
}
