package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func CreateProduct(context *gin.Context) (bool, error) {
	// Validate input
	var input model.CreateProductInput
	if err := context.ShouldBindJSON(&input); err != nil {
		return false, errors.New(err.Error())
	}

	product := model.Product{Name: input.Name, Description: input.Description}
	if err := DB.Create(&product).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
