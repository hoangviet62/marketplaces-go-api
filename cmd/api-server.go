package cmd

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
