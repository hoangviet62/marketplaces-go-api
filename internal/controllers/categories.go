package internal

import (
	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/categories"
	"net/http"
)

// func GetCategories(context *gin.Context) {
// 	var categories, pagination = service.GetCategories(context)
// 	context.JSON(http.StatusOK, gin.H{"data": categories, "pagination": pagination})
// }

func CreateCategory(context *gin.Context) {
	// Validate input

	category, err := service.CreateCategory(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": category})
}
