package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func SkuRoutes(router *gin.Engine) {
	router.GET("/skus", controllers.GetSkus)
	router.POST("/skus", controllers.CreateSku)
	router.GET("/skus/:id", controllers.GetSkuById)
	router.PATCH("/skus/:id", controllers.UpdateSku)
	router.DELETE("/skus/:id", controllers.DeleteSku)
}
