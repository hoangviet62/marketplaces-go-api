package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetOrderById(context *gin.Context) (model.Order, error) {
	var order model.Order
	if err := helpers.DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&order).Error; err != nil {
		return order, errors.New("record not found")
	}

	return order, nil
}
