package internal

import (
	"errors"
	// "fmt"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"strconv"
)

func UpdateBanner(context *gin.Context, attachments []model.Attachment) (model.Banner, error) {
	var banner model.Banner
	if err := DB.Where("id = ?", context.Param("id")).First(&banner).Error; err != nil {
		return banner, errors.New("Record not found")
	}

	// Validate input
	description := context.PostForm("description")
	linkTo := context.PostForm("link_to")
	priority, _ := strconv.ParseUint(context.PostForm("priority"), 10, 8)
	input := model.Banner{Description: description, LinkTo: linkTo, Priority: uint(priority), Attachments: attachments,}

	if err := DB.Model(&banner).Updates(input).Error; err != nil {
		return banner, errors.New(err.Error())
	}

	return banner, nil
}
