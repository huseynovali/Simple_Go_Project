package db

import (
	"log"
	"todo-app/config"
	"todo-app/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres(cfg *config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = DB.AutoMigrate(&model.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}
}
