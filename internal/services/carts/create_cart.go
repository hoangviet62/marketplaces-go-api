package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/users"
)

func CreateCart(context *gin.Context) (model.Cart, error) {
	user, err := service.GetCurrentUser(context)
	var cart model.Cart

	if err != nil {
		return cart, errors.New(err.Error())
	}

	cart = model.Cart{UserID: user.ID}

	if err := helpers.DB.Create(&cart).Error; err != nil {
		return cart, errors.New(err.Error())
	}

	return cart, nil
}
