package internal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	// "github.com/hoangviet62/marketplaces-go-api/cmd"
	"net/http"

	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

var DB *gorm.DB

func GetProducts(c *gin.Context) {
	var products []model.Product
	DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}
