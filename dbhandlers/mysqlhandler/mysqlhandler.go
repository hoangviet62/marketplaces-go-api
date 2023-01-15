package mysqlhandler

import (
	"fmt"

	"example.com/m/v2/envhandler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func conn() *gorm.DB {
	Host := envhandler.GetEnv("MYSQL_HOST")
	Port := envhandler.GetEnv("MYSQL_PORT")
	Username := envhandler.GetEnv("MYSQL_USERNAME")
	Password := envhandler.GetEnv("MYSQL_PASSWORD")
	Database := envhandler.GetEnv("MYSQL_DATABASE")
	Server := Host + ":" + Port
	dsn := Username + ":" + Password + "@tcp(" + Server + ")/" + Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("ERROR: ", err)
	}
	return db
}
