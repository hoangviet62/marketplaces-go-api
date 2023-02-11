package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"strconv"
)

func CreateSku(context *gin.Context) (model.Sku, error) {
	// Validate input
	// var input model.CreateProductInput
	description := context.PostForm("description")
	quantity, _ := strconv.Atoi(context.PostForm("quantity"))
	price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	status, _ := strconv.Atoi(context.PostForm("status"))
	productId, _ := strconv.ParseInt(context.PostForm("product_id"), 10, 64)
	sku := model.Sku{Description: description, ProductID: int32(productId), Quantity: quantity, Price: price, Status: status}

	if err := DB.Create(&sku).Error; err != nil {
		return sku, errors.New(err.Error())
	}

	return sku, nil
}
