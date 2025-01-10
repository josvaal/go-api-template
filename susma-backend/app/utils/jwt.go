package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/josvaal/susma-backend/app/models"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

func GenerateToken(account models.Account) (string, error) {
	println(secretKey)
	claims := jwt.MapClaims{
		"data": account,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (models.Account, error) {
	var account models.Account

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return account, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data, ok := claims["data"].(map[string]interface{})
		if !ok {
			return account, errors.New("datos inválidos en el token")
		}

		account.ID = int64(data["id"].(float64))
		account.Name = data["name"].(string)
		account.Lastname = data["lastname"].(string)
		account.Email = data["email"].(string)

		return account, nil
	}

	return account, errors.New("token inválido o expirado")
}
