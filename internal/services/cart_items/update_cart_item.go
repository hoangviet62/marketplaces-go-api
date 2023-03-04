package internal

import (
	"errors"
	// "fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"strconv"
)

func UpdateCartItem(context *gin.Context) (model.CartItem, error) {
	var cartItem model.CartItem
	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&cartItem).Error; err != nil {
		return cartItem, errors.New("record not found")
	}

	// Validate input
	price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	quantity, _ := strconv.ParseUint(context.PostForm("quantity"), 10, 8)

	input := model.CartItem{
		Price: price,
		Quantity: uint(quantity),
	}

	if err := helpers.DB.Model(&cartItem).Updates(input).Error; err != nil {
		return cartItem, errors.New(err.Error())
	}

	return cartItem, nil
}
