package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type DataJWT struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Email string `json:"email"`
}

var APPLICATION_NAME = "POS Mini"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret key")

func GenerateJWT(name string, email string) (string, error) {
	claims := DataJWT{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Name:  name,
		Email: email,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
