package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	cartService "github.com/hoangviet62/marketplaces-go-api/internal/services/carts"
	orderService "github.com/hoangviet62/marketplaces-go-api/internal/services/order_items"
	userService "github.com/hoangviet62/marketplaces-go-api/internal/services/users"
)

func CreateOrder(context *gin.Context) (model.Order, error) {
	user, err := userService.GetCurrentUser(context)
	var order model.Order
	if err != nil {
		return order, errors.New(err.Error())
	}

	tx := helpers.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return order, err
	}

	status := context.PostForm("status")
	code := context.PostForm("code")

	order = model.Order{UserID: user.ID, Status: status, Code: code}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return order, errors.New(err.Error())
	}

	for _, cartItem := range user.Cart.CartItem {
		orderService.CreateOrderItemTransaction(context, tx, order, cartItem)
	}

	_, deleteCartError := cartService.DeleteCartTransaction(context, tx, user.Cart.ID)

	if deleteCartError != nil {
		return order, errors.New(deleteCartError.Error())
	}

	tx.Commit()
	return order, nil
}
