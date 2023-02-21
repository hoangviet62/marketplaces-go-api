package jwt

import (
	"errors"
	"fmt"
	"gorm.io/gorm/clause"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	kong "github.com/hoangviet62/marketplaces-go-api/internal/services/kong"
)

func Decode(authToken string) (user model.User, err error) {
	var username string
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		username = fmt.Sprint(claims["username"])
		jwt := kong.FetchConsumerJwt(username)
		return []byte(jwt.Data[0].Secret), nil
	})

	_, ok := token.Claims.(jwt.MapClaims)
	record := model.User{}

	if ok && err == nil {
		record = model.User{Username: username}
		helpers.DB.Preload(clause.Associations).First(&record, "username = ?", record.Username)
		return record, nil
	}

	return record, errors.New(err.Error())
}
