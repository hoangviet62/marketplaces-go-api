package internal

import (
	"errors"
	// "fmt"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"strconv"
)

func UpdateSku(context *gin.Context, attachments []model.Attachment) (model.Sku, error) {
	var sku model.Sku
	if err := DB.Where("id = ?", context.Param("id")).First(&sku).Error; err != nil {
		return sku, errors.New("Record not found")
	}

	// Validate input
	description := context.PostForm("description")
	quantity, _ := strconv.Atoi(context.PostForm("quantity"))
	price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	status, _ := strconv.Atoi(context.PostForm("status"))
	productId, _ := strconv.ParseInt(context.PostForm("product_id"), 10, 64)

	input := model.Sku{
		Description: description,
		Quantity:    quantity,
		Price:       price,
		Status:      status,
		ProductID:   int32(productId),
		Attachments: attachments,
	}

	if err := DB.Model(&sku).Updates(input).Error; err != nil {
		return sku, errors.New(err.Error())
	}

	return sku, nil
}
