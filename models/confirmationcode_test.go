package model_test

import (
	"github.com/stretchr/testify/assert"
	model "github.com/trello-analog/backend/models"
	"testing"
	"time"
)

// confirmation code is expired
func TestConfirmationCode_IsCodeExpired(t *testing.T) {
	code := model.ConfirmationCode{
		ID:      0,
		Code:    "",
		UserId:  0,
		Expired: time.Now().Add(time.Hour * (-1)).UTC().String(),
	}

	assert.Equal(t, code.IsCodeExpired(), true)
}
