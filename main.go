package main

import (
	_ "go_learn/docs"
	"go_learn/internal/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// snake_game.RunGame()
	engine := router.SetupRouter()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":8000")
}
