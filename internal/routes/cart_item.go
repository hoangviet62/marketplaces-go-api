package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func CartItemRoutes(router *gin.Engine) {
	router.GET("/cart_items", controllers.GetCartItems)
	router.POST("/cart_items", controllers.CreateCartItem)
	router.GET("/cart_items/:id", controllers.GetCartItemById)
	router.PATCH("/cart_items/:id", controllers.UpdateCartItem)
	router.DELETE("/cart_items/:id", controllers.DeleteCartItem)
}
