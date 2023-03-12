package helpers

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySqlConnection() {
	user := viper.GetString("MYSQL.USERNAME")
	password := viper.GetString("MYSQL.PASSWORD")
	database := viper.GetString("MYSQL.DATABASE")
	host := viper.GetString("MYSQL.HOST")
	port := viper.GetString("MYSQL.PORT")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		log.Errorf(dsn)
		log.Error("MYSQL: ", err)
	}

	DB = db
}
