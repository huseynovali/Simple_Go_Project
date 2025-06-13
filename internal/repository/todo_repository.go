package repository

import (
	"todo-app/internal/db"
	"todo-app/internal/model"
)

func GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	result := db.DB.Find(&todos)
	return todos, result.Error
}

func CreateTodo(todo *model.Todo) error {
	return db.DB.Create(todo).Error
}
