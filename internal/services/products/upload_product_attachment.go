package internal

import (
	"errors"

	"github.com/gin-gonic/gin"
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"gorm.io/gorm/clause"
)

func GetProductAttachments(context *gin.Context) (map[string][]string, error) {
	// Validate input
	// var input model.CreateProductInput

	form, _ := context.MultipartForm()
	images := form.File["images"]
	medias := form.File["medias"]

	imagesPath, imageErr := helpers.UploadFile(context, images, "images", "products")

	if imageErr != nil {
		return nil, errors.New(imageErr.Error())
	}

	mediasPath, mediaErr := helpers.UploadFile(context, medias, "medias", "products")

	if mediaErr != nil {
		return nil, errors.New(mediaErr.Error())
	}

	attachments := map[string][]string{
		"images": imagesPath,
		"medias": mediasPath,
	}

	return attachments, nil
}

func UploadProductAttachment(productId uint, attachments []model.Attachment, attachment_type string) (model.Product, error) {
	var product model.Product

	if err := helpers.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		return product, errors.New("record not found")
	}

	if err := helpers.DB.Model(&product).Update("Attachments", attachments).Error; err != nil {
		return product, errors.New(err.Error())
	}

	if attachment_type == "medias" {
		if err := helpers.DB.Model(&product).Update("Medias", attachments).Error; err != nil {
			return product, errors.New(err.Error())
		}
	} else {
		if err := helpers.DB.Model(&product).Update("Images", attachments).Error; err != nil {
			return product, errors.New(err.Error())
		}
	}

	helpers.DB.Preload(clause.Associations).Where("id = ?", productId).First(&product)

	return product, nil
}
