package internal

import (
	"errors"
	// "fmt"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
	// "strconv"
)

func UpdateCategory(context *gin.Context) (model.Category, error) {
	category := model.Category{}
	if err := DB.Where("id = ?", context.Param("id")).First(&category).Error; err != nil {
		return category, errors.New("Record not found")
	}

	// Validate input
	name := context.PostForm("name")

	deleteImages := context.PostFormArray("delete_images")
	deleteMedias := context.PostFormArray("delete_medias")

	if len(deleteImages) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteImages)
		if error != nil {
			return category, errors.New(error.Error())
		}
	}

	if len(deleteMedias) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteMedias)
		if error != nil {
			return category, errors.New(error.Error())
		}
	}

	input := model.Category{
		Name: name,
	}

	if err := DB.Model(&category).Updates(input).Error; err != nil {
		return category, errors.New(err.Error())
	}

	return category, nil
}
