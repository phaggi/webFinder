package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) AuthenticateUser(username, password string) (string, error) {
	// Dummy authentication logic
	if username == "admin" && password == "password" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return "", err
		}

		return tokenString, nil
	}

	return "", errors.New("invalid credentials")
}
