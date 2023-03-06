package internal

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	jwtService "github.com/hoangviet62/marketplaces-go-api/internal/services/jwt"
	log "github.com/sirupsen/logrus"
)

func GetCurrentUser(context *gin.Context) (model.User, error) {
	token := strings.ReplaceAll(context.GetHeader("Authorization"), "Bearer ", "")
	user, err := jwtService.Decode(token)

	if err != nil {
		log.Error("[err] ", err)
		return user, errors.New("Bad Request")
	}

	return user, err
}
