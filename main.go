package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/hoangviet62/marketplaces-go-api/cmd"
	"github.com/hoangviet62/marketplaces-go-api/helpers"
	. "github.com/hoangviet62/marketplaces-go-api/helpers"
	model "github.com/hoangviet62/marketplaces-go-api/internal/models"
)

func main() {
	helpers.InitMySqlConnection()
	cmd.StartMigration()
	// Hard
	skuHard := model.SkuHard{ID: 1, CreatedAt: time.Now()}
	DB.Create(&skuHard)
	specHard := model.SpecHard{ID: 1, CreatedAt: time.Now(), SkuHardId: skuHard.ID}
	DB.Create(&specHard)
	log.Info("Hello1")
	if err := DB.Select("SpecHard").Delete(&skuHard).Error; err != nil {
		log.Info("Hello")
		log.Info(err)
	}
	cmd.StartApiServer()
}
