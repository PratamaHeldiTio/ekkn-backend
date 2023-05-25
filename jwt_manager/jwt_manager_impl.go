package jwtmanager

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtManagerImpl struct{}

func NewJwtManager() JwtManager {
	return &JwtManagerImpl{}
}

// func for generate jwt_manager jwt
func (manager *JwtManagerImpl) GenerateJwt(nim, role, profile string) (string, error) {
	// create peyload jwt_manager
	exp := 60 * time.Minute
	if role == "admin" {
		exp = 6 * time.Hour
	}

	claim := jwt.MapClaims{
		"id":      nim,
		"exp":     time.Now().Add(exp).Unix(),
		"role":    role,
		"profile": profile,
	}

	// generate jwt_manager
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// add sigin to jwt_manager
	key := []byte(os.Getenv("SECRET_JWT"))
	signedToken, err := token.SignedString(key)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// validate jwt_manager
func (manager *JwtManagerImpl) ValidateJwt(tokenStr string) (*jwt.Token, error) {
	// parse jwt_manager
	validToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid jwt manager")
		}

		key := []byte(os.Getenv("SECRET_JWT"))
		return []byte(key), nil
	})

	if err != nil {
		return validToken, err
	}

	return validToken, nil
}
