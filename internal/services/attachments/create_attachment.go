package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func CreateAttachment(context *gin.Context, attachment_type string, attachment_id uint, url string) (model.Attachment, error) {
	attachment := model.Attachment{AttachmentType: attachment_type, Url: url, AttachmentID: attachment_id}

	if err := helpers.DB.Create(&attachment).Error; err != nil {
		return attachment, errors.New(err.Error())
	}

	return attachment, nil
}
