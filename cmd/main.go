package main

import (
	"todo-app/config"
	"todo-app/internal/db"
	"todo-app/internal/router"
	"log"
	"github.com/joho/godotenv"

)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()
	db.InitPostgres(cfg)

	r := router.SetupRouter()
	r.Run(":8081")
}
