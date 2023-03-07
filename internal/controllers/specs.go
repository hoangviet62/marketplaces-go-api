package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/specs"
)

func GetSpecs(context *gin.Context) {
	var specs, pagination = service.GetSpecs(context)
	context.JSON(http.StatusOK, gin.H{"data": specs, "pagination": pagination})
}

func CreateSpec(context *gin.Context) {
	// Validate input
	spec, err := service.CreateSpec(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": spec})
}

func GetSpecById(context *gin.Context) {
	spec, err := service.GetSpecById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": spec})
}

func UpdateSpec(context *gin.Context) {
	updatedSpec, err := service.UpdateSpec(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedSpec})
}

func DeleteSpec(context *gin.Context) {
	isDeleted, err := service.DeleteSpec(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
