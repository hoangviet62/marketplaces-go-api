package main

import (
	"example.com/m/v2/envhandler"
)

func main() {
	Host := envhandler.GetEnv("MYSQL_HOST")
	Port := envhandler.GetEnv("MYSQL_PORT")
	Username := envhandler.GetEnv("MYSQL_USERNAME")
	Password := envhandler.GetEnv("MYSQL_PASSWORD")
	Database := envhandler.GetEnv("MYSQL_DATABASE")
}
