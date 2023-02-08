package internal

import (
	"github.com/gin-gonic/gin"
	attachmentService "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/products"
	"net/http"
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

	attachmentURLs, attachmentURLsErr := service.UploadProductAttachment(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	for _, imageURL := range imageURLs {
		attachmentService.CreateAttachment(context, "product_images", product.ID, imageURL)
	}

	for _, mediaURL := range mediaURLs {
		attachmentService.CreateAttachment(context, "product_medias", product.ID, mediaURL)
	}

	context.JSON(http.StatusCreated, gin.H{"data": product})
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
	updatedProduct, err := service.UpdateProduct(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
