package cmd

import (
	log "github.com/sirupsen/logrus"
)

func StartMigration() {
	log.Info("==== STARTING MIGRATIONS ====")
	log.Info(DB)
	log.Info("==== END MIGRATIONS ====")
}
