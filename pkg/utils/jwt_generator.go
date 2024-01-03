package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
)

func GenerateNewToken(u models.User) (string, error) {
	var secretKey = []byte("your-secret-key")
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return t, nil
}
