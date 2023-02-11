package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func Encode(user model.User, key string, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"iat":      time.Now().UTC().Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
