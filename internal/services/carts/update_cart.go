package internal

import (
	"errors"
	// "fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func UpdateCart(context *gin.Context) (model.Cart, error) {
	var cart model.Cart
	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&cart).Error; err != nil {
		return cart, errors.New("record not found")
	}

	input := model.Cart{}

	if err := helpers.DB.Model(&cart).Updates(input).Error; err != nil {
		return cart, errors.New(err.Error())
	}

	return cart, nil
}
