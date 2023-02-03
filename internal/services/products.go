package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func GetProducts() []model.Product {
	var products []model.Product
	DB.Find(&products)
	return products
}

func GetProductById(context *gin.Context) (model.Product, error) {
	var product model.Product
	if err := DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return product, errors.New("Record not found")
	}

	return product, nil
}

func CreateProduct(context *gin.Context) (bool, error) {
	// Validate input
	var input model.CreateProductInput
	if err := context.ShouldBindJSON(&input); err != nil {
		return false, errors.New(err.Error())
	}

	product := model.Product{Name: input.Name, Description: input.Description}
	if err := DB.Create(&product).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}

func UpdateProduct(context *gin.Context) (model.Product, error) {
	var product model.Product
	if err := DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return product, errors.New("Record not found")
	}

	// Validate input
	var input model.UpdateProductInput
	if err := context.ShouldBindJSON(&input); err != nil {
		return product, errors.New(err.Error())
	}

	if err := DB.Updates(input).Error; err != nil {
		return product, errors.New(err.Error())
	}

	DB.Model(&product).Updates(input)
	return product, nil
}

func DeleteProduct(context *gin.Context) (bool, error) {
	var product model.Product

	if err := DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := DB.Delete(&product).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
