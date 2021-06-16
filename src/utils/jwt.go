package utils

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type DataJWT struct {
	jwt.StandardClaims
	Name         string `json:"name"`
	Email        string `json:"email"`
	MerchantUuid string `json:"merchant"`
}

var APPLICATION_NAME = "POS Mini"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(os.Getenv("API_JWT_SECRET"))

func GenerateJWT(name string, email string, merchantUuid string) (string, error) {
	claims := DataJWT{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Name:         name,
		Email:        email,
		MerchantUuid: merchantUuid,
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

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(JWT_SIGNATURE_KEY), nil
	})
}
