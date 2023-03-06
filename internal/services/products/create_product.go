package internal

import (
	"errors"
	// "fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func CreateProduct(context *gin.Context) (model.Product, error) {
	// Validate input
	// var input model.CreateProductInput
	name := context.PostForm("name")
	description := context.PostForm("description")
	tag := context.PostForm("tag")
	isFeatured, _ := strconv.ParseBool(context.PostForm("is_featured"))
	status := model.ProductStatus(context.PostForm("status"))

	categoryId, _ := strconv.ParseUint(context.PostForm("category_id"), 10, 8)
	product := model.Product{Name: name, Tag: tag, Description: description, CategoryID: uint(categoryId), IsFeatured: isFeatured, Status: status}

	if err := helpers.DB.Create(&product).Error; err != nil {
		return product, errors.New(err.Error())
	}

	price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	quantity, _ := strconv.ParseUint(context.PostForm("quantity"), 10, 8)

	sku := model.Sku{Description: description, Quantity: uint(quantity), Price: price, ProductID: product.ID}

	if err := helpers.DB.Create(&sku).Error; err != nil {
		return product, errors.New(err.Error())
	}

	return product, nil
}
