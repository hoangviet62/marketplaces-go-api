package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/cart_items"
)

func GetCartItems(context *gin.Context) {
	var cartItems, pagination = service.GetCartItems(context)
	context.JSON(http.StatusOK, gin.H{"data": cartItems, "pagination": pagination})
}

func CreateCartItem(context *gin.Context) {
	// Validate input
	cartItem, err := service.CreateCartItem(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": cartItem})
}

func GetCartItemById(context *gin.Context) {
	cartItem, err := service.GetCartItemById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": cartItem})
}

func UpdateCartItem(context *gin.Context) {
	updatedCartItem, err := service.UpdateCartItem(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCartItem})
}

func DeleteCartItem(context *gin.Context) {
	isDeleted, err := service.DeleteCartItem(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
