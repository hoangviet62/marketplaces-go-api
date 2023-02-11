package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetSkuAttachments(context *gin.Context) (map[string][]string, error) {
	// Validate input
	// var input model.CreateProductInput

	form, _ := context.MultipartForm()
	images := form.File["images"]
	medias := form.File["medias"]

	imagesPath, imageErr := UploadFile(context, images, "images", "skus")

	if imageErr != nil {
		return nil, errors.New(imageErr.Error())
	}

	mediasPath, mediaErr := UploadFile(context, medias, "medias", "skus")

	if mediaErr != nil {
		return nil, errors.New(mediaErr.Error())
	}

	attachments := map[string][]string{
		"images": imagesPath,
		"medias": mediasPath,
	}

	return attachments, nil
}

func UploadSkuAttachment(skuId int32, attachments []model.Attachment, attachment_type string) (model.Sku, error) {
	var sku model.Sku

	if err := DB.Where("id = ?", skuId).First(&sku).Error; err != nil {
		return sku, errors.New("Record not found")
	}

	if err := DB.Model(&sku).Update("Attachments", attachments).Error; err != nil {
		return sku, errors.New(err.Error())
	}

	if attachment_type == "medias" {
		if err := DB.Model(&sku).Update("Medias", attachments).Error; err != nil {
			return sku, errors.New(err.Error())
		}
	} else {
		if err := DB.Model(&sku).Update("Images", attachments).Error; err != nil {
			return sku, errors.New(err.Error())
		}
	}

	DB.Preload(clause.Associations).Where("id = ?", skuId).First(&sku)

	return sku, nil
}
