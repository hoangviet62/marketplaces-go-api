package cmd

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	workDir, err := os.Getwd()
	cobra.CheckErr(err)
	viper.AddConfigPath(filepath.Join(workDir, "config"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")
	viper.Set("DEBUG", true)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func InitMySqlConnection() {
	user := viper.GetString("MYSQL.USERNAME")
	password := viper.GetString("MYSQL.PASSWORD")
	database := viper.GetString("MYSQL.DATABASE")
	host := viper.GetString("MYSQL.HOST")
	port := viper.GetString("MYSQL.PORT")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf(dsn)
		log.Error("MYSQL: ", err)
	}

	DB = db
}
