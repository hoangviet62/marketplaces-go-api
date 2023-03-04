package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func CartRoutes(router *gin.Engine) {
	router.GET("/carts", controllers.GetCarts)
	router.POST("/carts", controllers.CreateCart)
	router.GET("/carts/:id", controllers.GetCartById)
	router.PATCH("/carts/:id", controllers.UpdateCart)
	router.DELETE("/carts/:id", controllers.DeleteCart)
}
