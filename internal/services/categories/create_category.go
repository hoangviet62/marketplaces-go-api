package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func CreateCategory(context *gin.Context) (model.Category, error) {
	name := context.PostForm("name")
	category := model.Category{Name: name}

	if err := DB.Create(&category).Error; err != nil {
		return category, errors.New(err.Error())
	}

	return category, nil
}
