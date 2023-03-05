package internal

import (
	"errors"
	"fmt"

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

func GetCartItemByProductID(context *gin.Context, product_id uint) (model.CartItem, error) {
	fmt.Println("Product ID ", product_id)
	var cartItem model.CartItem
	if err := helpers.DB.Preload(clause.Associations).Where("product_id = ?", product_id).First(&cartItem).Error; err != nil {
		return cartItem, errors.New("record not found")
	}

	return cartItem, nil
}
