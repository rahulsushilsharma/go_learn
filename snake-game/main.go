package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type gamestate struct {
	board         [][]rune
	snakeLength   int
	snakePosition []int
	direction     int
	bound_x       int
	bound_y       int
}

const board_x, board_y = 48, 12

func updatePosition(state *gamestate) {

	switch state.direction {
	case 1:
		if state.snakePosition[0] > state.bound_y {
			state.snakePosition[0] = 0
		} else {
			state.snakePosition[0] += 1
		}
	case 2:
		if state.snakePosition[0] <= 0 {
			state.snakePosition[0] = state.bound_y - 1
		} else {
			state.snakePosition[0] -= 1
		}
	case 3:
		if state.snakePosition[1] > state.bound_x {
			state.snakePosition[1] = 0
		} else {
			state.snakePosition[1] += 1
		}
	case 4:
		if state.snakePosition[1] <= 0 {
			state.snakePosition[1] = state.bound_x - 1
		} else {
			state.snakePosition[1] -= 1
		}

	}

	// println(state.snakePosition[0], state.snakePosition[1])
}

func renderSnake(state *gamestate) {

	for i := range state.board {
		for j := range state.board[i] {
			state.board[i][j] = ' '
			if j == 0 || j == len(state.board[i])-1 {
				state.board[i][j] = '|'
			}
			if i == state.snakePosition[1] && j == state.snakePosition[0] {
				state.board[i][j] = '*'
			}
		}
	}

}

func getUserInput(inputChanel chan rune) {
	var input rune

	fmt.Scan(&input)
	println("usrInput ", input)
}

func gameLoop(state *gamestate, screen tcell.Screen, style tcell.Style, userInput rune) {

	// getUserInput(CurrentState)

	updatePosition(state)
	// renderSnake(CurrentState)

	width, height := screen.Size()

	for j := range height {
		screen.SetContent(0, j, '|', nil, style)
		screen.SetContent(width-1, j, '|', nil, style)
	}

	screen.SetContent(state.snakePosition[1], state.snakePosition[0], '*', nil, style)
}

func initGameState(screen tcell.Screen) gamestate {
	w, h := screen.Size()
	state := gamestate{
		board:         make([][]rune, w),
		snakeLength:   1,
		snakePosition: []int{3, 4},
		direction:     4,
		bound_x:       w,
		bound_y:       h,
	}

	return state

}

func readInput(inputChan chan rune, s tcell.Screen) {
	// reader := bufio.NewReader(os.Stdin)
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {

		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			inputChan <- ev.Rune()
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}

func updateDiraction(input rune, state *gamestate) {
	switch input {
	case 's':
		state.direction = 1
	case 'w':
		state.direction = 2
	case 'd':
		state.direction = 3
	case 'a':
		state.direction = 4
	}
}

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	inputChan := make(chan rune)
	quitChan := make(chan struct{})
	go readInput(inputChan, s)

	i, j := 0, 0

	width, height := s.Size()

	state := initGameState(s)
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	go func() {

		for {

			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				inputChan <- ev.Rune()
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					quit()
					close(quitChan)
				}
			}
		}

	}()
	userMessage := "hello user"
	for {
		s.Clear()
		select {
		case <-quitChan:
			return
		case input := <-inputChan:
			updateDiraction(input, &state)
			putString(s, 10, 10, defStyle, "key press")
		default:
			gameLoop(&state, s, tcell.StyleDefault, ' ')

			s.Show()
			i++
			j++
			putString(s, width-len(userMessage), height-10, defStyle, "hello")

			time.Sleep(100 * time.Millisecond)

		}

	}

	// Clear screen

}

func putString(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for i, r := range str {
		s.SetContent(x+i, y, r, nil, style)
	}
}
