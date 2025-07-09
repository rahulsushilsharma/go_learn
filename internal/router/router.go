package router

import (
	"go_learn/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()

	engine.GET("/users", handler.GetUsers)
	engine.POST("/users", handler.CreateUser)
	return engine
}
