package internal

// import (
// 	"errors"
// 	"github.com/gin-gonic/gin"
// 	. "github.com/hoangviet62/marketplaces-go-api/helpers"
// 	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
// )

// func GetCategoryById(context *gin.Context) (model.Category, error) {
// 	var category model.Category
// 	if err := DB.Preload("clause.Associations").Where("id = ?", context.Param("id")).First(&category).Error; err != nil {
// 		return category, errors.New("Record not found")
// 	}

// 	return category, nil
// }
