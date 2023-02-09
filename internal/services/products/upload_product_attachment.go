package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func GetProductAttachments(context *gin.Context) (map[string][]string, error) {
	// Validate input
	// var input model.CreateProductInput

	form, _ := context.MultipartForm()
	images := form.File["images"]
	medias := form.File["medias"]

	imagesPath, imageErr := UploadFile(context, images, "images", "products")

	if imageErr != nil {
		return nil, errors.New(imageErr.Error())
	}

	mediasPath, mediaErr := UploadFile(context, medias, "medias", "products")

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

	if err := DB.Where("id = ?", productId).First(&product).Error; err != nil {
		return product, errors.New("Record not found")
	}

	if err := DB.Model(&product).Update("Attachments", attachments).Error; err != nil {
		return product, errors.New(err.Error())
	}

	if attachment_type == "medias" {
		if err := DB.Model(&product).Update("Medias", attachments).Error; err != nil {
			return product, errors.New(err.Error())
		}
	} else {
		if err := DB.Model(&product).Update("Images", attachments).Error; err != nil {
			return product, errors.New(err.Error())
		}
	}

	DB.Preload("Attachments").Preload("Images").Preload("Medias").Where("id = ?", productId).First(&product)

	return product, nil
}
