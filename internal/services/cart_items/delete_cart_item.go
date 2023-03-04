package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteCartItem(context *gin.Context) (bool, error) {
	var cartItem model.CartItem

	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&cartItem).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := helpers.DB.Delete(&cartItem).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
