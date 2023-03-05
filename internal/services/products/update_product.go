package internal

import (
	"encoding/json"
	"errors"

	// "fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
)

func UpdateProduct(context *gin.Context) (model.Product, error) {
	product := model.Product{}
	if err := helpers.DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		return product, errors.New("record not found")
	}

	// Validate input
	name := context.PostForm("name")
	description := context.PostForm("description")
	tag := context.PostForm("tag")
	categoryId, _ := strconv.ParseUint(context.PostForm("category_id"), 10, 8)

	deleteImages := context.PostForm("delete_images")
	deleteMedias := context.PostForm("delete_medias")

	var deleteImagesArr []int
	_ = json.Unmarshal([]byte(deleteImages), &deleteImagesArr)

	var deleteMediasArr []int
	_ = json.Unmarshal([]byte(deleteMedias), &deleteMediasArr)

	if len(deleteImages) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteImagesArr)
		if error != nil {
			return product, errors.New(error.Error())
		}
		// for _, image := range deleteImages {
		// 	imageId, _ := strconv.ParseUint(image, 10, 8)
		// 	_, error := service.DeleteAttachment(context, uint(imageId))
		// 	if error != nil {
		// 		return product, errors.New(error.Error())
		// 	}
		// }
	}

	if len(deleteMedias) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteMediasArr)
		if error != nil {
			return product, errors.New(error.Error())
		}
		// for _, image := range deleteMedias {
		// 	imageId, _ := strconv.ParseUint(image, 10, 8)
		// 	_, error := service.DeleteAttachment(context, uint(imageId))
		// if error != nil {
		// 	return product, errors.New(error.Error())
		// }
		// }
	}

	input := model.Product{
		Name:        name,
		Description: description,
		CategoryID:  uint(categoryId),
		Tag:         tag,
	}

	if err := helpers.DB.Model(&product).Updates(input).Error; err != nil {
		return product, errors.New(err.Error())
	}

	return product, nil
}
