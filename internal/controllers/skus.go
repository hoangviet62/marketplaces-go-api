package internal

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	attachmentService "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/skus"
)

func GetSkus(context *gin.Context) {
	var skus, pagination = service.GetSkus(context)
	context.JSON(http.StatusOK, gin.H{"data": skus, "pagination": pagination})
}

func CreateSku(context *gin.Context) {
	// Validate input

	sku, err := service.CreateSku(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentURLs, attachmentURLsErr := service.GetSkuAttachments(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	var images []model.Attachment
	var medias []model.Attachment

	for _, imageURL := range imageURLs {
		attachment, err := attachmentService.CreateAttachment(context, "sku_images", sku.ID, imageURL)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		images = append(images, attachment)
	}

	for _, mediaURL := range mediaURLs {
		attachment, err := attachmentService.CreateAttachment(context, "sku_medias", sku.ID, mediaURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		medias = append(medias, attachment)
	}

	var updatedSku model.Sku
	if len(images) > 0 {
		updatedSku, err = service.UploadSkuAttachment(sku.ID, images, "images")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if len(medias) > 0 {
		updatedSku, err = service.UploadSkuAttachment(sku.ID, medias, "medias")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	result := sku

	if !reflect.DeepEqual(updatedSku, model.Sku{}) {
		result = updatedSku
	}

	context.JSON(http.StatusCreated, gin.H{"data": result})
}

func GetSkuById(context *gin.Context) {
	sku, err := service.GetSkuById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": sku})
}

func UpdateSku(context *gin.Context) {
	updatedSku, err := service.UpdateSku(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedSku})
}

func DeleteSku(context *gin.Context) {
	isDeleted, err := service.DeleteSku(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
