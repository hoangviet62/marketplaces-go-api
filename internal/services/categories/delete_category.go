package internal

// import (
// 	"errors"
// 	"github.com/gin-gonic/gin"
// 	. "github.com/hoangviet62/marketplaces-go-api/helpers"
// 	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
// )

// func DeleteCategory(context *gin.Context) (bool, error) {
// 	var category model.Category

// 	if err := DB.Where("id = ?", context.Param("id")).First(&category).Error; err != nil {
// 		return false, errors.New(err.Error())
// 	}

// 	if err := DB.Delete(&category).Error; err != nil {
// 		return false, errors.New(err.Error())
// 	}

// 	return true, nil
// }
