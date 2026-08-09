package utils

import (
	"gofiber-restful-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"expired_at": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(config.JWTSecret())
}
