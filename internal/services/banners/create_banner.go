package internal

import (
	"strconv"
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func CreateBanner(context *gin.Context) (model.Banner, error) {
	description := context.PostForm("description")
	linkTo := context.PostForm("link_to")
	priority, _ := strconv.ParseUint(context.PostForm("priority"), 10, 8)

	banner := model.Banner{Description: description, LinkTo: linkTo, Priority: uint(priority)}

	if err := DB.Create(&banner).Error; err != nil {
		return banner, errors.New(err.Error())
	}

	return banner, nil
}
