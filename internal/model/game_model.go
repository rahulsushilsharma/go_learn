package model

type Gamestate struct {
	// Board         [][]rune `json:"board"`
	SnakeLength   int     `json:"snakeLength"`
	SnakePosition [][]int `json:"SnakePosition"`
	Direction     int     `json:"Direction"`
	Bound_x       int     `json:"Bound_x"`
	Bound_y       int     `json:"Bound_y"`
	Food_x        int     `json:"Food_x"`
	Food_y        int     `json:"Food_y"`
	Game_over     bool
}
