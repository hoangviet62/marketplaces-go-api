package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func OrderItemRoutes(router *gin.Engine) {
	router.GET("/order_items", controllers.GetOrderItems)
	// router.POST("/order_items", controllers.CreateOrderItem)
	router.GET("/order_items/:id", controllers.GetOrderItemById)
	// router.PATCH("/order_items/:id", controllers.UpdateOrderItem)
	// router.DELETE("/order_items/:id", controllers.DeleteOrderItem)
}
