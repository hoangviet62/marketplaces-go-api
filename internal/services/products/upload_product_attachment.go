package internal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
)

func UploadProductAttachment(context *gin.Context) (map[string][]string, error) {
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

	fmt.Println(attachments)
	// attachments := append(imagesPath, mediasPath...)

	return attachments, nil
}
