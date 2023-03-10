package internal

import (
	"errors"
	// "fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func UpdateOrder(context *gin.Context) (model.Order, error) {
	var order model.Order
	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&order).Error; err != nil {
		return order, errors.New("record not found")
	}

	code := context.PostForm("code")
	status := context.PostForm("status")

	input := model.Order{Code: code, Status: status}

	if err := helpers.DB.Model(&order).Updates(input).Error; err != nil {
		return order, errors.New(err.Error())
	}

	return order, nil
}
