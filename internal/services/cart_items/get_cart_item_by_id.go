package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetCartItemById(context *gin.Context) (model.CartItem, error) {
	var cartItem model.CartItem
	if err := helpers.DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&cartItem).Error; err != nil {
		return cartItem, errors.New("record not found")
	}

	return cartItem, nil
}
