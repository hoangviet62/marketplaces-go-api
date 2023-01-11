package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/envhandler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port := envhandler.GetEnv("PORT")
	ServerAndPort := ":" + Port

	fmt.Printf("Server started at %s", ServerAndPort)
	http.ListenAndServe(ServerAndPort, nil)
}
