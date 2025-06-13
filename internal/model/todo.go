package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title" binding:"required" gorm:"not null"`
	Completed bool   `json:"completed"`
}
