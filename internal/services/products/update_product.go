package internal

import (
	"errors"
	// "fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	// "strconv"
)

func UpdateProduct(context *gin.Context, attachments []model.Attachment) (model.Product, error) {
	var product model.Product
	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return product, errors.New("record not found")
	}

	// Validate input
	name := context.PostForm("name")
	description := context.PostForm("description")

	input := model.Product{
		Name:        name,
		Description: description,
		Attachments: attachments,
	}

	if err := helpers.DB.Model(&product).Updates(input).Error; err != nil {
		return product, errors.New(err.Error())
	}

	return product, nil
}
