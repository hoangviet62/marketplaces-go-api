package internal

import (
	"github.com/gin-gonic/gin"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm"
)

func CreateOrderItemTransaction(context *gin.Context, tx *gorm.DB, order model.Order, cartItem ...model.CartItem) (model.OrderItem, error) {
	orderItem := model.OrderItem{Quantity: cartItem[0].Quantity, Price: cartItem[0].Price, OrderID: order.ID, ProductID: cartItem[0].Product.ID, SkuID: cartItem[0].SkuID}
	if err := tx.Create(&orderItem).Error; err != nil {
		tx.Rollback()
		return orderItem, err
	}

	return orderItem, nil
}
