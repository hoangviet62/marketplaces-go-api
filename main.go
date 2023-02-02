package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/cmd"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
)

func main() {
	r := gin.Default()

	helpers.InitMySqlConnection()
	cmd.StartMigration()
	cmd.StartApiServer()

	r.Run()
}
