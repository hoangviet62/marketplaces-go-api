package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func CreateCategory(context *gin.Context) (model.Category, error) {
	var input model.CreateCategoryInput
	category := model.Category{Name: input.Name}
	if err := context.ShouldBindJSON(&input); err != nil {
		return category, errors.New(err.Error())
	}

	if err := DB.Create(&category).Error; err != nil {
		return category, errors.New(err.Error())
	}

	return category, nil
}
