package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	RegisterTodoRoutes(api.Group("/todos"))

	return r
}
