package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetProducts)
}
