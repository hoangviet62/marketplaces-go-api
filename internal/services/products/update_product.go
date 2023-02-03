package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func UpdateProduct(context *gin.Context) (model.Product, error) {
	var product model.Product
	if err := DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return product, errors.New("Record not found")
	}

	// Validate input
	var input model.UpdateProductInput
	if err := context.ShouldBindJSON(&input); err != nil {
		return product, errors.New(err.Error())
	}

	if err := DB.Updates(input).Error; err != nil {
		return product, errors.New(err.Error())
	}

	DB.Model(&product).Updates(input)
	return product, nil
}
