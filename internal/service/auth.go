package service

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func (as *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			log.Fatalf("ValidateToken error: %v", isvalid)
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(tokenString), nil
	})
}
