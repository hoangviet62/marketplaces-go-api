package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetProductById(context *gin.Context) (model.Product, error) {
	var product model.Product
	if err := helpers.DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return product, errors.New("record not found")
	}

	return product, nil
}
