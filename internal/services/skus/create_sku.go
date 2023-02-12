package internal

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func CreateSku(context *gin.Context) (model.Sku, error) {
	// Validate input
	// var input model.CreateProductInput
	description := context.PostForm("description")
	quantity, _ := strconv.ParseUint(context.PostForm("quantity"), 10, 8)
	price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	status, _ := strconv.ParseUint(context.PostForm("status"), 10, 8)
	productId, _ := strconv.ParseUint(context.PostForm("product_id"), 10, 8)
	sku := model.Sku{Description: description, ProductID: uint(productId), Quantity: uint(quantity), Price: price, Status: uint(status)}

	if err := DB.Create(&sku).Error; err != nil {
		return sku, errors.New(err.Error())
	}

	return sku, nil
}
