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

func GetTodoByID(id uint) (*model.Todo, error) {
	todo, err := repository.GetTodoByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateTodo(todo *model.Todo) error {
	return repository.UpdateTodo(todo)
}