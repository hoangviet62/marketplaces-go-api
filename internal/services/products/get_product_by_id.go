package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func GetProductById(context *gin.Context) (model.Product, error) {
	var product model.Product
	if err := DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return product, errors.New("Record not found")
	}

	return product, nil
}
