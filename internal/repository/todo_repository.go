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

func UpdateTodo(todo *model.Todo) error {
	return db.DB.Save(todo).Error
}

func GetTodoByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	result := db.DB.First(&todo, "id = ?", id) // For primary key lookup
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}
