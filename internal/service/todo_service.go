package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

func GetTodos() ([]model.Todo, error) {
	return repository.GetAllTodos()
}

func AddTodo(todo *model.Todo) error {
	return repository.CreateTodo(todo)
}
