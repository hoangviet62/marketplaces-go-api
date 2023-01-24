package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/cmd"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func GetCategories(c *gin.Context) {
	categories := []model.Category{}
	cmd.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}