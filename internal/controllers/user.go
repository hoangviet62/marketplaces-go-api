package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/users"
)

func CreateUser(context *gin.Context) {
	created, user, err := service.CreateUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{
		"status": created,
		"user":   user,
	}
	context.JSON(http.StatusCreated, gin.H{"data": data})
}

func GetCurrentUser(context *gin.Context) {
	user, err := service.GetCurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": user})
}