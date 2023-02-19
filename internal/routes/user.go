package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users/register", controllers.CreateUser)
	router.GET("/users/me", controllers.GetCurrentUser)
}
