package internal

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	attachmentService "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/categories"
)

func GetCategories(context *gin.Context) {
	var categories, pagination = service.GetCategories(context)
	context.JSON(http.StatusOK, gin.H{"data": categories, "pagination": pagination})
}

func CreateCategory(context *gin.Context) {
	// Validate input

	category, err := service.CreateCategory(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentURLs, attachmentURLsErr := service.GetCategoryAttachments(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	var images []model.Attachment
	var medias []model.Attachment

	for _, imageURL := range imageURLs {
		attachment, err := attachmentService.CreateAttachment(context, "category_images", category.ID, imageURL)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		images = append(images, attachment)
	}

	for _, mediaURL := range mediaURLs {
		attachment, err := attachmentService.CreateAttachment(context, "category_medias", category.ID, mediaURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		medias = append(medias, attachment)
	}

	var updatedCategory model.Category
	if len(images) > 0 {
		updatedCategory, err = service.UploadCategoryAttachment(category.ID, images, "images")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if len(medias) > 0 {
		updatedCategory, err = service.UploadCategoryAttachment(category.ID, medias, "medias")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	result := category

	if !reflect.DeepEqual(updatedCategory, model.Category{}) {
		result = updatedCategory
	}

	context.JSON(http.StatusCreated, gin.H{"data": result})
}

func GetCategoryById(context *gin.Context) {
	category, err := service.GetCategoryById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": category})
}

func UpdateCategory(context *gin.Context) {
	category, err := service.UpdateCategory(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentURLs, attachmentURLsErr := service.GetCategoryAttachments(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	var images []model.Attachment
	var medias []model.Attachment

	for _, imageURL := range imageURLs {
		attachment, err := attachmentService.CreateAttachment(context, "category_images", category.ID, imageURL)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		images = append(images, attachment)
	}

	for _, mediaURL := range mediaURLs {
		attachment, err := attachmentService.CreateAttachment(context, "category_medias", category.ID, mediaURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		medias = append(medias, attachment)
	}

	var updatedCategory model.Category
	if len(images) > 0 {
		updatedCategory, err = service.UploadCategoryAttachment(category.ID, images, "images")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if len(medias) > 0 {
		updatedCategory, err = service.UploadCategoryAttachment(category.ID, medias, "medias")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	result := category

	if !reflect.DeepEqual(updatedCategory, model.Product{}) {
		result = updatedCategory
	}

	context.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteCategory(context *gin.Context) {
	isDeleted, err := service.DeleteCategory(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
