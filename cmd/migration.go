package cmd

import (
	helpers "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
	log "github.com/sirupsen/logrus"
)

func StartMigration() {
	log.Info("==== STARTING MIGRATIONS ====")
	helpers.DB.AutoMigrate(
		// &model.Attachment{},
		// &model.Banner{},
		// &model.CartItem{},
		// &model.Cart{},
		// &model.Category{},
		// &model.Country{},
		// &model.OrderItem{},
		// &model.Order{},
		// &model.Product{},
		// &model.Sku{},
		// &model.Spec{},
		// &model.User{},
		&model.SkuHard{},
		&model.SpecHard{},
	)
	log.Info("==== END MIGRATIONS ====")
}
