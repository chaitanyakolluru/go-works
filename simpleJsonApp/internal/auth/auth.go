package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var adminMap = map[string]string{"admin": "password"}

func getHmacSecret() string {
	hmacSecret, isExist := os.LookupEnv("HMAC_SECRET")
	if !isExist {
		return ""
	}

	return hmacSecret
}

func GenerateToken() (string, error) {
	hmacSecret := getHmacSecret()
	if hmacSecret == "" {
		return "", errors.New("environment variable HMAC_SECRET not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"admin": adminMap["admin"]})
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseAndValidate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return getHmacSecret(), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["admin"] == adminMap["admin"], nil
	}

	return false, nil
}
