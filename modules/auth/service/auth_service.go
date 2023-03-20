package service

import "github.com/dgrijalva/jwt-go"

type AuthService interface {
	GenerateTokenJwt(nim, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
