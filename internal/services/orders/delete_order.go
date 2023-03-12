package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteOrder(context *gin.Context) (bool, error) {
	var order model.Order

	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&order).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := helpers.DB.Delete(&order).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
