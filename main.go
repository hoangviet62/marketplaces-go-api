package main

import "github.com/hoangviet62/marketplaces-go-api/cmd"

func main() {
	cmd.InitMySqlConnection()
	cmd.StartMigration()
	cmd.StartApiServer()
}
