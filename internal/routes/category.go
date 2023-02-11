package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func CategoryRoutes(router *gin.Engine) {
	router.GET("/categories", controllers.GetCategories)
	router.POST("/categories", controllers.CreateCategory)
	router.GET("/categories/:id", controllers.GetCategoryById)
	router.PATCH("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
}
