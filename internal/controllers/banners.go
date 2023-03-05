package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	attachmentService "github.com/hoangviet62/marketplaces-go-api/internal/services/attachments"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/banners"
)

func GetBanners(context *gin.Context) {
	var banners, pagination = service.GetBanners(context)
	context.JSON(http.StatusOK, gin.H{"data": banners, "pagination": pagination})
}

func CreateBanner(context *gin.Context) {
	// Validate input

	banner, err := service.CreateBanner(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentURLs, attachmentURLsErr := service.GetBannerAttachments(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	var images []model.Attachment
	var medias []model.Attachment

	for _, imageURL := range imageURLs {
		attachment, err := attachmentService.CreateAttachment(context, "banner_images", banner.ID, imageURL)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		images = append(images, attachment)
	}

	for _, mediaURL := range mediaURLs {
		attachment, err := attachmentService.CreateAttachment(context, "banner_medias", banner.ID, mediaURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		medias = append(medias, attachment)
	}

	var updatedBanner model.Banner
	if images != nil {
		updatedBanner, err = service.UploadBannerAttachment(banner.ID, images, "images")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if medias != nil {
		updatedBanner, err = service.UploadBannerAttachment(banner.ID, medias, "medias")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusCreated, gin.H{"data": updatedBanner})
}

func GetBannerById(context *gin.Context) {
	banner, err := service.GetBannerById(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": banner})
}

func UpdateBanner(context *gin.Context) {
	banner, err := service.UpdateBanner(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentURLs, attachmentURLsErr := service.GetBannerAttachments(context)

	if attachmentURLsErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": attachmentURLsErr.Error()})
		return
	}

	imageURLs := attachmentURLs["images"]
	mediaURLs := attachmentURLs["medias"]

	var images []model.Attachment
	var medias []model.Attachment

	for _, imageURL := range imageURLs {
		attachment, err := attachmentService.CreateAttachment(context, "banner_images", banner.ID, imageURL)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		images = append(images, attachment)
	}

	for _, mediaURL := range mediaURLs {
		attachment, err := attachmentService.CreateAttachment(context, "banner_medias", banner.ID, mediaURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		medias = append(medias, attachment)
	}

	var updatedBanner model.Banner
	if images != nil {
		updatedBanner, err = service.UploadBannerAttachment(banner.ID, images, "images")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if medias != nil {
		updatedBanner, err = service.UploadBannerAttachment(banner.ID, medias, "medias")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedBanner})
}

func DeleteBanner(context *gin.Context) {
	isDeleted, err := service.DeleteBanner(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": isDeleted})
}
