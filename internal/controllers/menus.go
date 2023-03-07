package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/hoangviet62/marketplaces-go-api/internal/services/menu"
)

func GetMenus(context *gin.Context) {
	var menus = service.Customer()
	context.JSON(http.StatusOK, gin.H{"data": menus})
}
