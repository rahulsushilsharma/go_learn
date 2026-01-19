package main

import (
	_ "go_learn/docs"
	imagetoterminal "go_learn/internal/image_to_terminal"
	"go_learn/internal/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve() {
	engine := router.SetupRouter()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":8000")
}

func main() {
	// go Serve()

	// snake_game.RunGame()
	imagetoterminal.RunParse()
}
