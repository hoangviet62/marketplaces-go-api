package envhandler

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(keyToFetch string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(keyToFetch)
}
