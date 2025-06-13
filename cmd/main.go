package main

import (
	"todo-app/config"
	"todo-app/internal/db"
	"todo-app/internal/router"
)

func main() {
	cfg := config.LoadConfig()
	db.InitPostgres(cfg)

	r := router.SetupRouter()
	r.Run(":8080")
}
