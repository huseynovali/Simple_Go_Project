package router

import (
	"todo-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	RegisterTodoRoutes(api.Group("/todos").Use(middleware.AuthMiddleware()).(*gin.RouterGroup))
	RegisterAuthRoutes(api.Group("/auth"))
	return r
}
