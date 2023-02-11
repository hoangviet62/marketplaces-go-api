package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetSkuById(context *gin.Context) (model.Sku, error) {
	var sku model.Sku
	if err := DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&sku).Error; err != nil {
		return sku, errors.New("Record not found")
	}

	return sku, nil
}
