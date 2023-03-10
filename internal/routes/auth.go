package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/auth/sign_in", controllers.SignIn)
	router.DELETE("/auth/sign_out", controllers.SignOut)
}
