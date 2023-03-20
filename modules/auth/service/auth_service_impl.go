package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthServiceImpl struct{}

var SECRET_KEY = []byte("INI SECRET KEY BUAT PEMBELAJARAN")

func NewAuthServiceImpl() AuthService {
	return &AuthServiceImpl{}
}

// func for generate token jwt
func (service *AuthServiceImpl) GenerateTokenJwt(nim, role string) (string, error) {
	// create peyload token
	claim := jwt.MapClaims{
		"nim":  nim,
		"exp":  time.Now().Add(60 * time.Minute).Unix(),
		"role": role,
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// add sigin to token
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// validate token
func (service *AuthServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	// parse token
	validToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return validToken, err
	}

	return validToken, nil
}
