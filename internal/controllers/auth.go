package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/auth"
)

func SignIn(context *gin.Context) {
	user := service.SignIn(context)
	httpStatus := http.StatusUnauthorized
	if user.Status {
		httpStatus = http.StatusAccepted
	}
	context.JSON(httpStatus, gin.H{"data": user})
}

func SignOut(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"data": ""})
}
