package internal

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hoangviet62/marketplaces-go-api/internal/controllers"
)

func BannerRoutes(router *gin.Engine) {
	router.GET("/banners", controllers.GetBanners)
	router.POST("/banners", controllers.CreateBanner)
	router.GET("/banners/:id", controllers.GetBannerById)
	router.PATCH("/banners/:id", controllers.UpdateBanner)
	router.DELETE("/banners/:id", controllers.DeleteBanner)
}
