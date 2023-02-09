package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// file_type = "medias" or "images"
// Directory = "products", "categories", etc.
func UploadFile(context *gin.Context, files []*multipart.FileHeader, file_type string, directory string) ([]string, error) {
	var filesPath []string

	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		replacer := strings.NewReplacer(" ", "-", ".", "-")
		filename := replacer.Replace(strings.ToLower(originalFileName)) + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

		out, err := os.Create("public/" + directory + "/" + file_type + "/" + filename)
		if err != nil {
			return nil, err
		}

		defer out.Close()

		fileExtract, err := file.Open()

		_, err = io.Copy(out, fileExtract)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		filePath := "/public/" + directory + "/" + file_type + "/" + filename
		filesPath = append(filesPath, filePath)
	}

	return filesPath, nil
}
