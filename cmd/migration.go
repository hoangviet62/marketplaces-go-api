package cmd

import (
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	log "github.com/sirupsen/logrus"
)

func StartMigration() {
	log.Info("==== STARTING MIGRATIONS ====")
	DB.AutoMigrate(&model.User{})
	log.Info("==== END MIGRATIONS ====")
}
