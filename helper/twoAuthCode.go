package helper

import (
	"math/rand"
	"strconv"
)

func GenerateTwoAuthCode() string {
	randomNumber := rand.Int63()
	return string([]rune(strconv.FormatInt(randomNumber, 10))[0:4])
}
