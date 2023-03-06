package main

import (
	"github.com/hoangviet62/marketplaces-go-api/cmd"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	menu "github.com/hoangviet62/marketplaces-go-api/internal/services/menu"
)

func main() {
	helpers.InitMySqlConnection()
	// cmd.StartMigration()
	menu.Customer()
	cmd.StartApiServer()
}
