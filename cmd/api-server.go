package cmd

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

	initMySqlConnection()
}

func StartApiServer() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(time.Second * 60))
	router.Use(middleware.SetHeader("content-type", "application/json"))

	log.Infof("Start listening on 0.0.0.0:%s", viper.GetString("PORT"))
	server := "0.0.0.0:" + viper.GetString("PORT")
	err := http.ListenAndServe(server, router)
	if err != nil {
		panic(err)
	}
}

func initMySqlConnection() {
	user := viper.GetString("MYSQL.USERNAME")
	password := viper.GetString("MYSQL.PASSWORD")
	database := viper.GetString("MYSQL.DATABASE")
	host := viper.GetString("MYSQL.HOST")
	port := viper.GetString("MYSQL.PORT")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf(dsn)
		log.Error("MYSQL: ", err)
	}
}
