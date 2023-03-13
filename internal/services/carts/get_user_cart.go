package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/users"
	"gorm.io/gorm/clause"
)

func GetUserCart(context *gin.Context) (model.Cart, error) {
	user, err := service.GetCurrentUser(context)
	cart := model.Cart{}

	if err != nil {
		return cart, errors.New(err.Error())
	}

	if err := helpers.DB.Preload(clause.Associations).Preload("CartItem.Product.Images").Preload("CartItem.Product.Medias").FirstOrCreate(&cart, model.Cart{UserID: user.ID}).Error; err != nil {
		return cart, errors.New("record not found")
	}

	return cart, nil
}
