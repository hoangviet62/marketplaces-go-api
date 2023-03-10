package internal

import (
	"errors"
	// "fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func UpdateSku(context *gin.Context) (model.Sku, error) {
	var sku model.Sku
	if err := DB.Where("id = ?", context.Param("id")).First(&sku).Error; err != nil {
		return sku, errors.New("Record not found")
	}

	// Validate input
	description := context.PostForm("description")
	quantity, _ := strconv.Atoi(context.PostForm("quantity"))
	price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	status, _ := strconv.Atoi(context.PostForm("status"))
	productId, _ := strconv.ParseUint(context.PostForm("product_id"), 10, 32)

	input := model.Sku{
		Description: description,
		Quantity:    uint(quantity),
		Price:       price,
		Status:      uint(status),
		ProductID:   uint(productId),
	}

	if err := DB.Model(&sku).Updates(input).Error; err != nil {
		return sku, errors.New(err.Error())
	}

	return sku, nil
}
