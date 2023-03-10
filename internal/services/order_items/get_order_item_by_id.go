package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetOrderItemById(context *gin.Context) (model.OrderItem, error) {
	var orderItem model.OrderItem
	if err := helpers.DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&orderItem).Error; err != nil {
		return orderItem, errors.New("record not found")
	}

	return orderItem, nil
}
