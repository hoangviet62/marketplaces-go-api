package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/cmd"
	"github.com/hoangviet62/marketplaces-go-api/internal/routes"
)

func main() {
	r := gin.Default()

	cmd.InitMySqlConnection()
	cmd.StartMigration()
	cmd.StartApiServer()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	routes.CategoryRoutes(r)

	r.Run()
}
