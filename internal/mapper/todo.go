// internal/mapper/todo_mapper.go
package mapper

import (
	"todo-app/internal/dto/response"
	"todo-app/internal/model"
)

func ToTodoResponse(todo model.Todo) response.TodoResponse {
	return response.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt.Format("2006-01-02 15:04:05")}
}
