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

func GetTodoByID(id string) (response.TodoResponse, error) {
	todo, err := repository.GetTodoByID(id)
	if err != nil {
		return response.TodoResponse{}, err
	}

   var todoResponse response.TodoResponse
	todoResponse = mapper.ToTodoResponse(*todo)



	return todoResponse, nil
}

func UpdateTodo(todo *model.Todo) error {
	return repository.UpdateTodo(todo)
}
