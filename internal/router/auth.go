
package router

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handler"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", handler.Register)
	rg.POST("/login", handler.Login)
}
