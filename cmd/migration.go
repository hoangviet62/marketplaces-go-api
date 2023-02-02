package cmd

import (
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	log "github.com/sirupsen/logrus"
)

func StartMigration() {
	log.Info("==== STARTING MIGRATIONS ====")
	DB.AutoMigrate(
		&model.User{},
		&model.Permission{},
		&model.Role{},
	)
	log.Info("==== END MIGRATIONS ====")
}
