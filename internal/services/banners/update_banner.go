package internal

import (
	"encoding/json"
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

	deleteImages := context.PostForm("delete_images")
	deleteMedias := context.PostForm("delete_medias")

	var deleteImagesArr []int
	_ = json.Unmarshal([]byte(deleteImages), &deleteImagesArr)

	var deleteMediasArr []int
	_ = json.Unmarshal([]byte(deleteMedias), &deleteMediasArr)

	if len(deleteImages) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteImagesArr)
		if error != nil {
			return banner, errors.New(error.Error())
		}
	}

	if len(deleteMedias) > 0 {
		_, error := service.DeleteAttachmentsArr(context, deleteMediasArr)
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
