package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func Decode(user model.User, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, _ := token.SignedString(secret)

	return tokenString
}
