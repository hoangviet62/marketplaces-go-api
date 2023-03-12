package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetSpecById(context *gin.Context) (model.Spec, error) {
	var spec model.Spec
	if err := DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&spec).Error; err != nil {
		return spec, errors.New("Record not found")
	}

	return spec, nil
}
