package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangviet62/marketplaces-go-api/cmd"
)

func main() {
	r := gin.Default()

	cmd.InitMySqlConnection()
	cmd.StartMigration()
	cmd.StartApiServer()

	r.Run()
}
