package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/carts"
)

func GetCarts(context *gin.Context) {
	var carts, pagination = service.GetCarts(context)
	context.JSON(http.StatusOK, gin.H{"data": carts, "pagination": pagination})
}

func CreateCart(context *gin.Context) {
	// Validate input
	cart, err := service.CreateCart(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": cart})
}

func GetCartById(context *gin.Context) {
	cart, err := service.GetCartById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": cart})
}

func UpdateCart(context *gin.Context) {
	updatedCart, err := service.UpdateCart(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCart})
}

func DeleteCart(context *gin.Context) {
	isDeleted, err := service.DeleteCart(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
