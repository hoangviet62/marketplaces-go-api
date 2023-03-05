package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	attachmentService "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/products"
)

func GetProducts(context *gin.Context) {
	var products, pagination = service.GetProducts(context)
	context.JSON(http.StatusOK, gin.H{"data": products, "pagination": pagination})
}

func CreateProduct(context *gin.Context) {
	// Validate input

	product, err := service.CreateProduct(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentURLs, attachmentURLsErr := service.GetProductAttachments(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	var images []model.Attachment
	var medias []model.Attachment

	for _, imageURL := range imageURLs {
		attachment, err := attachmentService.CreateAttachment(context, "product_images", product.ID, imageURL)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		images = append(images, attachment)
	}

	for _, mediaURL := range mediaURLs {
		attachment, err := attachmentService.CreateAttachment(context, "product_medias", product.ID, mediaURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		medias = append(medias, attachment)
	}

	var updatedProduct model.Product
	if images != nil {
		updatedProduct, err = service.UploadProductAttachment(product.ID, images, "images")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if medias != nil {
		updatedProduct, err = service.UploadProductAttachment(product.ID, medias, "medias")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusCreated, gin.H{"data": updatedProduct})
}

func GetProductById(context *gin.Context) {
	product, err := service.GetProductById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(context *gin.Context) {
	product, err := service.UpdateProduct(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentURLs, attachmentURLsErr := service.GetProductAttachments(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	var images []model.Attachment
	var medias []model.Attachment

	for _, imageURL := range imageURLs {
		attachment, err := attachmentService.CreateAttachment(context, "product_images", product.ID, imageURL)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		images = append(images, attachment)
	}

	for _, mediaURL := range mediaURLs {
		attachment, err := attachmentService.CreateAttachment(context, "product_medias", product.ID, mediaURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		medias = append(medias, attachment)
	}

	var updatedProduct model.Product

	if images != nil {
		updatedProduct, err = service.UploadProductAttachment(product.ID, images, "images")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if medias != nil {
		updatedProduct, err = service.UploadProductAttachment(product.ID, medias, "medias")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedProduct})
}

func DeleteProduct(context *gin.Context) {
	isDeleted, err := service.DeleteProduct(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
