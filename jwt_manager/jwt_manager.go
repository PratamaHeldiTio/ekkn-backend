package jwtmanager

import "github.com/dgrijalva/jwt-go"

type JwtManager interface {
	GenerateJwt(nim, role string) (string, error)
	ValidateJwt(token string) (*jwt.Token, error)
}
