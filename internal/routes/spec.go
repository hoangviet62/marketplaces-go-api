package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func SpecRoutes(router *gin.Engine) {
	router.GET("/specs", controllers.GetSpecs)
	router.POST("/specs", controllers.CreateSpec)
	router.GET("/specs/:id", controllers.GetSpecById)
	router.PATCH("/specs/:id", controllers.UpdateSpec)
	router.DELETE("/specs/:id", controllers.DeleteSpec)
}
