package cmd

import (
	log "github.com/sirupsen/logrus"

	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func StartMigration() {
	log.Info("==== STARTING MIGRATIONS ====")
	DB.AutoMigrate(&model.Category{})
	log.Info("==== END MIGRATIONS ====")
}
