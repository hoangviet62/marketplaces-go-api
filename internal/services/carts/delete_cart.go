package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteCart(context *gin.Context) (bool, error) {
	var cart model.Cart

	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&cart).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := helpers.DB.Delete(&cart).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
