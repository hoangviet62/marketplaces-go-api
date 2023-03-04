package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func CreateCartItem(context *gin.Context) (model.CartItem, error) {
	var input model.CreateCartItemInput
	product := model.Product{}
	cart := model.Cart{}
	cartItem := model.CartItem{}

	if err := context.ShouldBindJSON(&input); err != nil {
		return cartItem, errors.New(err.Error())
	}

	helpers.DB.Preload(clause.Associations).First(&product, "id = ?", input.ProductID)
	helpers.DB.Preload(clause.Associations).First(&cart, "id = ?", input.CartID)
	// var cartItem model.CartItem
	// product := model.Product{}
	// var cart model.Cart

	// price, _ := strconv.ParseFloat(context.PostForm("price"), 64)
	// quantity, _ := strconv.ParseUint(context.PostForm("quantity"), 10, 8)
	// cart_id, _ := strconv.ParseUint(context.PostForm("cart_id"), 10, 8)

	// helpers.DB.Preload(clause.Associations).First(&product, "id = ?", context.PostForm("product_id"))
	// helpers.DB.Preload(clause.Associations).First(&cart, "id = ?", cart_id)

	cartItem = model.CartItem{Product: product, Cart: cart, CartID: input.CartID, Price: input.Price, Quantity: input.Quantity}

	if err := helpers.DB.Create(&cartItem).Error; err != nil {
		return cartItem, errors.New(err.Error())
	}

	return cartItem, nil
}
