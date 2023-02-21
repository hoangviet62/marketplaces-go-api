package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetBannerById(context *gin.Context) (model.Banner, error) {
	var banner model.Banner
	if err := DB.Preload(clause.Associations).Where("id = ?", context.Param("id")).First(&banner).Error; err != nil {
		return banner, errors.New("Record not found")
	}

	return banner, nil
}
