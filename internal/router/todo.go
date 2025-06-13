package router

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handler"
)

func RegisterTodoRoutes(rg *gin.RouterGroup) {
	rg.GET("/", handler.GetTodos)
	rg.POST("/", handler.CreateTodo)
	rg.GET("/:id", handler.GetTodoByID)
	rg.PUT("/:id", handler.UpdateTodo)
	rg.DELETE("/:id", handler.DeleteTodo)
	rg.PUT("/:id/toggle", handler.ToggleTodoCompletion)
}
