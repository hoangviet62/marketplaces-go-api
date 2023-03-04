package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetCartById(context *gin.Context) (model.Cart, error) {
	var cart model.Cart
	if err := helpers.DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&cart).Error; err != nil {
		return cart, errors.New("record not found")
	}

	return cart, nil
}
