package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func CartRoutes(router *gin.Engine) {
	router.GET("/cart", controllers.GetCarts)
	router.POST("/cart", controllers.CreateCart)
	router.GET("/cart/:id", controllers.GetCartById)
	router.PATCH("/cart/:id", controllers.UpdateCart)
	router.DELETE("/cart/:id", controllers.DeleteCart)
}
