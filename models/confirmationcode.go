package model

import (
	uuid "github.com/nu7hatch/gouuid"
	"time"
)

type ConfirmationCode struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	UserId    int    `json:"user_id"`
	Expired   string `json:"expired"`
	Confirmed bool   `json:"confirmed"`
	Last      bool   `json:"last"`
}

func NewConfirmationCode(userId int) *ConfirmationCode {
	code, _ := uuid.NewV4()

	return &ConfirmationCode{
		Code:      code.String(),
		UserId:    userId,
		Expired:   time.Now().Add(time.Hour * 24).UTC().String(),
		Confirmed: false,
		Last:      true,
	}
}

func (cc *ConfirmationCode) IsCodeExpired() bool {
	now := time.Now().Unix()
	layout := "2006-01-02 15:04:05 -0700 MST"
	codeExpire, err := time.Parse(layout, cc.Expired)

	if err != nil {
		return false
	}

	return now > codeExpire.Unix()
}

func (cc *ConfirmationCode) MakeConfirmed() *ConfirmationCode {
	cc.Confirmed = true
	return cc
}

func (cc *ConfirmationCode) IsConfirmed() bool {
	return cc.Confirmed
}

func (cc *ConfirmationCode) IsLast() bool {
	return cc.Last
}

func (cc *ConfirmationCode) MakeIrrelevant() *ConfirmationCode {
	cc.Last = false
	return cc
}
