package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"strconv"
)

func CreateProduct(context *gin.Context) (model.Product, error) {
	// Validate input
	// var input model.CreateProductInput
	name := context.PostForm("name")
	description := context.PostForm("description")
	categoryId, _ := strconv.Atoi(context.PostForm("category_id"))
	product := model.Product{Name: name, Description: description, CategoryID: categoryId}

	if err := DB.Create(&product).Error; err != nil {
		return product, errors.New(err.Error())
	}

	return product, nil
}
