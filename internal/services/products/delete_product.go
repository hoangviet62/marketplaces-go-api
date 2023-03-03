package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteProduct(context *gin.Context) (bool, error) {
	product := model.Product{}

	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := helpers.DB.Delete(&product).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
