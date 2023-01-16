package cmd

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	Server := "0.0.0.0:" + viper.GetString("PORT")
	err := http.ListenAndServe(Server, router)
	if err != nil {
		panic(err)
	}
}
