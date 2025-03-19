package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var refreshSecret = []byte("refresh_secret")

func GenerateRefreshToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}
