package internal

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/datatypes"
)

func UpdateSpec(context *gin.Context) (model.Spec, error) {
	var spec model.Spec
	if err := DB.Where("id = ?", context.Param("id")).First(&spec).Error; err != nil {
		return spec, errors.New("Record not found")
	}

	// Validate input
	description, _ := context.GetPostForm("description")
	specDescription := datatypes.JSON([]byte(description))
	productId, _ := strconv.ParseUint(context.PostForm("product_id"), 10, 32)

	input := model.Spec{
		Description: specDescription,
		ProductID:   uint(productId),
	}

	if err := DB.Model(&spec).Updates(input).Error; err != nil {
		return spec, errors.New(err.Error())
	}

	return spec, nil
}
