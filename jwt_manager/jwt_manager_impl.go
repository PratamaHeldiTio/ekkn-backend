package jwtmanager

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtManagerImpl struct{}

var SECRET_KEY = []byte("INI SECRET KEY BUAT PEMBELAJARAN")

func NewJwtManagerImpl() JwtManager {
	return &JwtManagerImpl{}
}

// func for generate jwt_manager jwt
func (service *JwtManagerImpl) GenerateJwt(nim, role string) (string, error) {
	// create peyload jwt_manager
	claim := jwt.MapClaims{
		"nim":  nim,
		"exp":  time.Now().Add(60 * time.Minute).Unix(),
		"role": role,
	}

	// generate jwt_manager
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// add sigin to jwt_manager
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// validate jwt_manager
func (service *JwtManagerImpl) ValidateJwt(token string) (*jwt.Token, error) {
	// parse jwt_manager
	validToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid jwt_manager")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return validToken, err
	}

	return validToken, nil
}
