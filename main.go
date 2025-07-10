package main

import (
	_ "go_learn/docs"
	"go_learn/internal/router"
	"go_learn/internal/snake_game"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve() {
	engine := router.SetupRouter()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":8000")
}

func main() {
	go Serve()

	snake_game.RunGame()

}
