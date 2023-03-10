package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteCategory(context *gin.Context) (bool, error) {
	category := model.Category{}

	if err := DB.Preload("Products").Where("id = ?", context.Param("id")).First(&category).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := DB.Select("Images", "Medias", "Attachments", "Products").Delete(&category).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
