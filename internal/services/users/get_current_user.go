package internal

import (
	"github.com/gin-gonic/gin"
	"errors"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	jwtService "github.com/hoangviet62/marketplaces-go-api/internal/services/jwt"
)

func GetCurrentUser(context *gin.Context) (model.User, error) {
	token := context.GetHeader("Authorization")
	user, err := jwtService.Decode(token)

	if err != nil {
		return user, errors.New("Bad Request")
	}

	return user, err
}
