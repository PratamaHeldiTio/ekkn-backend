package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthServiceImpl struct{}

func NewAuthServiceImpl() AuthService {
	return &AuthServiceImpl{}
}

// func for generate token jwt
func (service *AuthServiceImpl) GenerateTokenJwt(nim string) (string, error) {
	// create peyload token
	claim := jwt.MapClaims{
		"nim": nim,
		"iat": time.Now().Unix(),
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// add sigin to token
	SECRET_KEY := []byte("INI SECRET KEY BUAT PEMBELAJARAN")
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
