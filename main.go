package main

// "github.com/hoangviet62/marketplaces-go-api/cmd"

import (
	"github.com/hoangviet62/marketplaces-go-api/cmd"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	kong "github.com/hoangviet62/marketplaces-go-api/internal/services/kong"
)

func main() {
	helpers.InitMySqlConnection()
	kong.CreateConsumer("viet")
	cmd.StartMigration()
	cmd.StartApiServer()
}
