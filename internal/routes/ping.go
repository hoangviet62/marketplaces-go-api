package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func PingRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
