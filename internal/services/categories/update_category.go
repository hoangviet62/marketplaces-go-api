package internal

// import (
// 	"errors"
// 	// "fmt"
// 	"github.com/gin-gonic/gin"
// 	. "github.com/hoangviet62/marketplaces-go-api/helpers"
// 	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
// 	// "strconv"
// )

// func UpdateCategory(context *gin.Context, attachments []model.Attachment) (model.Category, error) {
// 	var category model.Category
// 	if err := DB.Where("id = ?", context.Param("id")).First(&category).Error; err != nil {
// 		return category, errors.New("Record not found")
// 	}

// 	// Validate input
// 	name := context.PostForm("name")
// 	description := context.PostForm("description")

// 	input := model.Category{
// 		Name:        name,
// 		Description: description,
// 		Attachments: attachments,
// 	}

// 	if err := DB.Model(&category).Updates(input).Error; err != nil {
// 		return category, errors.New(err.Error())
// 	}

// 	return category, nil
// }
