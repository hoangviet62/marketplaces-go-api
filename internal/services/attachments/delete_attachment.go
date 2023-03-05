package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func DeleteAttachment(context *gin.Context, id uint) (bool, error) {
	attachment := model.Attachment{}

	if err := DB.Where("id = ?", id).First(&attachment).Error; err != nil {
		return false, errors.New(err.Error())
	}

	if err := DB.Delete(&attachment).Error; err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}

func DeleteAttachmentsArr(context *gin.Context, attachments []int) (bool, error) {
	for _, attachment := range attachments {
		_, error := DeleteAttachment(context, uint(attachment))
		if error != nil {
			return false, errors.New(error.Error())
		}
	}
	return true, nil
}
