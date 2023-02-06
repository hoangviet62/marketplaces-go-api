package internal

import (
	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/products"
	"net/http"
)

func GetProducts(context *gin.Context) {
	var products, pagination = service.GetProducts(context)
	context.JSON(http.StatusOK, gin.H{"data": products, "pagination": pagination})
}

func GetProductById(context *gin.Context) {
	product, err := service.GetProductById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(context *gin.Context) {
	// Validate input
	created, err := service.GetProductById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": created})
}

func UpdateProduct(context *gin.Context) {
	updatedProduct, err := service.GetProductById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedProduct})
}

func DeleteProduct(context *gin.Context) {
	isDeleted, err := service.DeleteProduct(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
