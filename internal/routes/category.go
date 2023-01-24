package routes

import (
	"github.com/hoangviet62/marketplaces-go-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	router.GET("/categories", controllers.GetCategories)
}