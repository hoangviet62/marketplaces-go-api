package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteSku(context *gin.Context) (bool, error) {
	var sku model.Sku

	if err := DB.Where("id = ?", context.Param("id")).First(&sku).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := DB.Delete(&sku).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
