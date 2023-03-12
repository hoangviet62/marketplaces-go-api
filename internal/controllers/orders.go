package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/orders"
)

func GetOrders(context *gin.Context) {
	var orders, pagination = service.GetOrders(context)
	context.JSON(http.StatusOK, gin.H{"data": orders, "pagination": pagination})
}

func CreateOrder(context *gin.Context) {
	// Validate input
	order, err := service.CreateOrder(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": order})
}

func GetOrderById(context *gin.Context) {
	order, err := service.GetOrderById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": order})
}

func UpdateOrder(context *gin.Context) {
	updatedOrder, err := service.UpdateOrder(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedOrder})
}

func DeleteOrder(context *gin.Context) {
	isDeleted, err := service.DeleteOrder(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
