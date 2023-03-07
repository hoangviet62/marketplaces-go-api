package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteSpec(context *gin.Context) (bool, error) {
	var spec model.Spec

	if err := DB.Where("id = ?", context.Param("id")).First(&spec).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := DB.Delete(&spec).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}
