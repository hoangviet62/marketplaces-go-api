package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"

	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func (c *BaseController) GetProducts(context *gin.Context) {
	var products []model.Product
	c.DB.Find(&products)
	context.JSON(http.StatusOK, gin.H{"data": products})
}

func (c *BaseController) GetProductById(context *gin.Context) {
	var product model.Product

	if err := c.DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": product})
}

func (c *BaseController) CreateProduct(context *gin.Context) {
	// Validate input
	var input model.CreateProductInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Product{Name: input.Name, Description: input.Description}
	c.DB.Create(&product)

	context.JSON(http.StatusOK, gin.H{"data": product})
}

func (c *BaseController) UpdateProduct(context *gin.Context) {
	var product model.Product
	if err := model.DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input model.UpdateProductInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.DB.Model(&product).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": product})
}

func (c *BaseController) DeleteProduct(context *gin.Context) {
	var product model.Product
	if err := model.DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.DB.Delete(&book)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
