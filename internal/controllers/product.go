package internal

import (
	"github.com/gin-gonic/gin"
	// "github.com/hoangviet62/marketplaces-go-api/cmd"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	"net/http"
)

func GetProducts(c *gin.Context) {
	var products []model.Product
	DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}
