package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	"gorm.io/gorm/clause"

	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func GetBannerAttachments(context *gin.Context) (map[string][]string, error) {
	// Validate input
	// var input model.CreateProductInput

	form, _ := context.MultipartForm()
	images := form.File["images"]
	medias := form.File["medias"]

	imagesPath, imageErr := UploadFile(context, images, "images", "banners")

	if imageErr != nil {
		return nil, errors.New(imageErr.Error())
	}

	mediasPath, mediaErr := UploadFile(context, medias, "medias", "banners")

	if mediaErr != nil {
		return nil, errors.New(mediaErr.Error())
	}

	attachments := map[string][]string{
		"images": imagesPath,
		"medias": mediasPath,
	}

	return attachments, nil
}

func UploadBannerAttachment(bannerId uint, attachments []model.Attachment, attachment_type string) (model.Banner, error) {
	var banner model.Banner

	if err := DB.Where("id = ?", bannerId).First(&banner).Error; err != nil {
		return banner, errors.New("Record not found")
	}

	if err := DB.Model(&banner).Update("Attachments", attachments).Error; err != nil {
		return banner, errors.New(err.Error())
	}

	if attachment_type == "medias" {
		if err := DB.Model(&banner).Update("Medias", attachments).Error; err != nil {
			return banner, errors.New(err.Error())
		}
	} else {
		if err := DB.Model(&banner).Update("Images", attachments).Error; err != nil {
			return banner, errors.New(err.Error())
		}
	}

	DB.Preload(clause.Associations).Where("id = ?", bannerId).First(&banner)

	return banner, nil
}
