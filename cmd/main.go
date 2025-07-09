package main

import (
	"go_learn/internal/router"
)

func main() {
	// snake_game.RunGame()
	engine := router.SetupRouter()
	engine.Run(":8000")
}
