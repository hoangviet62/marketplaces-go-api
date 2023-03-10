package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/order_items"
)

func GetOrderItems(context *gin.Context) {
	var orderItems, pagination = service.GetOrderItems(context)
	context.JSON(http.StatusOK, gin.H{"data": orderItems, "pagination": pagination})
}

func GetOrderItemById(context *gin.Context) {
	order, err := service.GetOrderItemById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": order})
}
