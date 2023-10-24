package main

import (
	"fmt"
	"log"
	"mongodb/config"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load file .env : %v", err)
	}

	conf := config.GetDBConfig()
	conn, err := config.Connect(conf)
	if err != nil {
		log.Fatalf("Error connect to database : %v", err)
		return
	}

	log.Printf("Server running at http://localhost:%v...", conf.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%v", conf.ServerPort), routeInit(conn))
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
