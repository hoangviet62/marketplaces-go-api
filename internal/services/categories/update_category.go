package internal

import (
	"encoding/json"
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

	deleteImages := context.PostForm("delete_images")
	deleteMedias := context.PostForm("delete_medias")

	var deleteImagesArr []int
	_ = json.Unmarshal([]byte(deleteImages), &deleteImagesArr)

	var deleteMediasArr []int
	_ = json.Unmarshal([]byte(deleteMedias), &deleteMediasArr)

	if len(deleteImages) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteImagesArr)
		if error != nil {
			return category, errors.New(error.Error())
		}
	}

	if len(deleteMedias) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteMediasArr)
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
