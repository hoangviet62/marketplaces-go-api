package internal

import (
	"errors"
	"strconv"
	// "fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func UpdateOrderItem(context *gin.Context) (model.OrderItem, error) {
	var orderItem model.OrderItem

	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&orderItem).Error; err != nil {
		return orderItem, errors.New("record not found")
	}

	price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	quantity, _ := strconv.ParseUint(context.PostForm("quantity"), 10, 8)
	productId, _ := strconv.ParseUint(context.PostForm("product_id"), 10, 8)

	input := model.OrderItem{
		Price:     price,
		Quantity:  uint(quantity),
		ProductID: uint(productId),
	}

	if err := helpers.DB.Model(&orderItem).Updates(input).Error; err != nil {
		return orderItem, errors.New(err.Error())
	}

	return orderItem, nil
}
