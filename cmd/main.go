package main

import (
	"todo-app/config"
	"todo-app/internal/db"
	"todo-app/internal/router"
	"log"
	"github.com/joho/godotenv"

)

func main() {

	err := godotenv.Load() // və ya godotenv.Load("../.env") əgər cmd/ içindənsə
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Konfiqurasiya və DB bağlantısı
	cfg := config.LoadConfig()
	db.InitPostgres(cfg)

	// Router-i başlat
	r := router.SetupRouter()
	r.Run(":8081")
}
