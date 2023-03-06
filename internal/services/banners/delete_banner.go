package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteBanner(context *gin.Context) (bool, error) {
	var banner model.Banner

	if err := DB.Where("id = ?", context.Param("id")).First(&banner).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := DB.Select("Images", "Medias", "Attachments").Delete(&banner).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
