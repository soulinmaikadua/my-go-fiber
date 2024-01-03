package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewToken() (string, error) {
	var secretKey = []byte("your-secret-key")
	claims := jwt.MapClaims{
		"name": "John Doe",
		"admin": true,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return t, nil
}