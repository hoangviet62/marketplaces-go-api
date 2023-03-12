package internal

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/datatypes"
)

func CreateSpec(context *gin.Context, specParams ...model.Spec) (model.Spec, error) {
	// Validate input
	// var input model.CreateProductInput
	spec := model.Spec{}
	if specParams != nil {
		spec = specParams[0]
	} else {
		description, _ := context.GetPostForm("description")
		specDescription := datatypes.JSON([]byte(description))

		productId, _ := strconv.ParseUint(context.PostForm("product_id"), 10, 8)
		spec = model.Spec{Description: specDescription, ProductID: uint(productId)}
	}

	if err := DB.Create(&spec).Error; err != nil {
		return spec, errors.New(err.Error())
	}

	return spec, nil
}
