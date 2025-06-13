package router

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/todos", handler.GetTodos)
		api.POST("/todos", handler.CreateTodo)
	}

	return r
}
