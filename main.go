package main

import (
	"github.com/hoangviet62/marketplaces-go-api/cmd"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
)

func main() {
	helpers.InitMySqlConnection()
	// cmd.StartMigration()
	cmd.StartApiServer()
}
