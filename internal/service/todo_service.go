package service

import (
	"todo-app/internal/dto/response"
	"todo-app/internal/mapper"
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

func GetTodos() ([]response.TodoResponse, error) {
	todos, err := repository.GetAllTodos()
	if err != nil {
		return nil, err
	}

	var todoResponses []response.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, mapper.ToTodoResponse(todo))
	}

	return todoResponses, nil
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
