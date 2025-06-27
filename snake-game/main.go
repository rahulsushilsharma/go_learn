package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

type gamestate struct {
	board         [][]rune
	snakeLength   int
	snakePosition []int
	direction     int
}

func updatePosition(state *gamestate) {
	bound_x := len(state.board[0])
	bound_y := len(state.board)

	switch state.direction {
	case 1:
		if state.snakePosition[0] >= bound_x {
			state.snakePosition[0] = 0
		} else {
			state.snakePosition[0] += 1
		}
	case 2:
		if state.snakePosition[0] <= 0 {
			state.snakePosition[0] = 0
		} else {
			state.snakePosition[0] -= 1
		}
	case 3:
		if state.snakePosition[1] >= bound_y {
			state.snakePosition[1] = 0
		} else {
			state.snakePosition[1] += 1
		}
	case 4:
		if state.snakePosition[1] <= 0 {
			state.snakePosition[1] = 0
		} else {
			state.snakePosition[1] -= 1
		}

	}
}

func renderSnake(state *gamestate) {
}

func gameLoop(CurrentState *gamestate) {
	// updatePosition(CurrentState)

	for _, v := range CurrentState.board {
		for _, char := range v {
			print(string(char))
		}
		println("")
	}

}

func main() {
	println("hello world")

	state := gamestate{
		board: [][]rune{
			{'X', 'O', 'X'},
			{'O', 'X', 'O'},
			{'X', 'O', 'X'},
		},
		snakePosition: []int{0, 0},
		snakeLength:   1,
		direction:     0,
	}

	for {
		gameLoop(&state)
		println("score: ", state.snakeLength)
		time.Sleep(1 * time.Second)
		CallClear()
	}

}
