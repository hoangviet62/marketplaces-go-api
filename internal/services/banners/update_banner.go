package internal

import (
	"errors"
	// "fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
)

func UpdateBanner(context *gin.Context) (model.Banner, error) {
	var banner model.Banner
	if err := DB.Where("id = ?", context.Param("id")).First(&banner).Error; err != nil {
		return banner, errors.New("Record not found")
	}

	// Validate input
	description := context.PostForm("description")
	linkTo := context.PostForm("link_to")
	priority, _ := strconv.ParseUint(context.PostForm("priority"), 10, 8)

	deleteImages := context.PostFormArray("delete_images")
	deleteMedias := context.PostFormArray("delete_medias")

	if len(deleteImages) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteImages)
		if error != nil {
			return banner, errors.New(error.Error())
		}
	}

	if len(deleteMedias) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteMedias)
		if error != nil {
			return banner, errors.New(error.Error())
		}
	}

	input := model.Banner{Description: description, LinkTo: linkTo, Priority: uint(priority)}

	if err := DB.Model(&banner).Updates(input).Error; err != nil {
		return banner, errors.New(err.Error())
	}

	return banner, nil
}
