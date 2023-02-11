package internal

// import (
// 	"errors"
// 	"github.com/gin-gonic/gin"
// 	. "github.com/hoangviet62/marketplaces-go-api/helpers"
// 	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
// )

// func GetCategoryAttachments(context *gin.Context) (map[string][]string, error) {
// 	// Validate input
// 	// var input model.CreateProductInput

// 	form, _ := context.MultipartForm()
// 	images := form.File["images"]
// 	medias := form.File["medias"]

// 	imagesPath, imageErr := UploadFile(context, images, "images", "products")

// 	if imageErr != nil {
// 		return nil, errors.New(imageErr.Error())
// 	}

// 	mediasPath, mediaErr := UploadFile(context, medias, "medias", "products")

// 	if mediaErr != nil {
// 		return nil, errors.New(mediaErr.Error())
// 	}

// 	attachments := map[string][]string{
// 		"images": imagesPath,
// 		"medias": mediasPath,
// 	}

// 	return attachments, nil
// }

// func UploadCategoryAttachment(categoryId uint, attachments []model.Attachment, attachment_type string) (model.Category, error) {
// 	var product model.Product

// 	if err := DB.Where("id = ?", categoryId).First(&category).Error; err != nil {
// 		return category, errors.New("Record not found")
// 	}

// 	if err := DB.Model(&category).Update("Attachments", attachments).Error; err != nil {
// 		return category, errors.New(err.Error())
// 	}

// 	if attachment_type == "medias" {
// 		if err := DB.Model(&category).Update("Medias", attachments).Error; err != nil {
// 			return category, errors.New(err.Error())
// 		}
// 	} else {
// 		if err := DB.Model(&category).Update("Images", attachments).Error; err != nil {
// 			return category, errors.New(err.Error())
// 		}
// 	}

// 	DB.Preload("clause.Associations").Where("id = ?", categoryId).First(&category)

// 	return category, nil
// }
