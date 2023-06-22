package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var adminMap = map[string]string{"admin": "password"}

func getHmacSecret() []byte {
	hmacSecret, isExist := os.LookupEnv("HMAC_SECRET")
	if !isExist {
		return []byte{}
	}

	return []byte(hmacSecret)
}

func GenerateToken() (string, error) {
	hmacSecret := getHmacSecret()
	if len(hmacSecret) == 0 {
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
	return true, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		tokenSlice := strings.Split(tokenValue, "Bearer ")

		var token string
		if len(tokenSlice) == 1 {
			tokenV, err := GenerateToken()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				log.Fatalf("fail to generate token, %v", err)
			}
			token = tokenV
		} else {
			token = tokenSlice[len(tokenSlice)-1]
		}

		valid, err := ParseAndValidate(token)

		fmt.Println("valid is", valid, err)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
