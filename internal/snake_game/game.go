package snake_game

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	"go_learn/internal/model"

	"github.com/gdamore/tcell/v2"
)

func updatePosition(state *model.Gamestate, index int) {
	checkEating(state)

	for i := len(state.SnakePosition) - 1; i >= 0; i-- {

		if i == 0 {
			continue
		}
		state.SnakePosition[i][2] = state.SnakePosition[i-1][2]
		state.SnakePosition[i][1] = state.SnakePosition[i-1][1]
		state.SnakePosition[i][0] = state.SnakePosition[i-1][0]

	}
	switch state.SnakePosition[0][2] {
	case 1:
		if state.SnakePosition[0][0] > state.Bound_y {
			state.SnakePosition[0][0] = 0
		} else {
			state.SnakePosition[0][0] += 1
		}
	case 2:
		if state.SnakePosition[0][0] <= 0 {
			state.SnakePosition[0][0] = state.Bound_y - 1
		} else {
			state.SnakePosition[0][0] -= 1
		}
	case 3:
		if state.SnakePosition[0][1] > state.Bound_x {
			state.SnakePosition[0][1] = 0
		} else {
			state.SnakePosition[0][1] += 1
		}
	case 4:
		if state.SnakePosition[0][1] <= 0 {
			state.SnakePosition[0][1] = state.Bound_x - 1
		} else {
			state.SnakePosition[0][1] -= 1
		}
	}
	checkEating(state)

}

func logFile(state *model.Gamestate) {

	file, err := os.Open("logfile.txt")
	if err != nil {
		panic(err)
	}

	for i := range state.SnakePosition {
		sentence := fmt.Sprintf("snake %d,%d,%d ", state.SnakePosition[i][0], state.SnakePosition[i][1], state.SnakePosition[i][2])
		_, err = file.Write([]byte(sentence))

		if err != nil {
			panic(err)
		}

	}

}

func gameLoop(state *model.Gamestate, screen tcell.Screen, style tcell.Style, index int) {
	if state.Game_over {
		putString(screen, state.Bound_x/2, state.Bound_y/2, style, "Game Over")
		putString(screen, state.Bound_x/2-7, state.Bound_y/2+2, style, "Press r to restart the game")

		return

	}
	updatePosition(state, index)
	screen.SetContent(state.Food_x, state.Food_y, '0', nil, style)

	for j := range state.Bound_y - 2 {
		screen.SetContent(0, j, '|', nil, style)
		screen.SetContent(state.Bound_x, j, '|', nil, style)
	}

	for i := range state.SnakePosition {
		draw_head := '>'

		switch state.SnakePosition[0][2] {
		case 1:
			draw_head = 'V'
		case 2:
			draw_head = '^'
		case 3:
			draw_head = '>'
		case 4:
			draw_head = '<'
		}

		if i == 0 {
			screen.SetContent(state.SnakePosition[i][1], state.SnakePosition[i][0], draw_head, nil, style)
			continue
		}

		screen.SetContent(state.SnakePosition[i][1], state.SnakePosition[i][0], '*', nil, style)
	}

}

func initGameState(screen tcell.Screen) model.Gamestate {
	w, h := screen.Size()
	state := model.Gamestate{
		SnakeLength:   1,
		SnakePosition: make([][]int, 2),
		Direction:     4,
		Bound_x:       w / 3,
		Bound_y:       h,
		Food_x:        10,
		Food_y:        10,
	}
	for i := range state.SnakePosition {
		state.SnakePosition[i] = []int{3, 4 + i, 1}
	}

	return state

}

func generateFood(state *model.Gamestate) (int, int) {
	var x, y int
	for {
		x = rand.IntN(state.Bound_x-2) + 1
		found := false
		for i := range state.SnakePosition {
			if state.SnakePosition[i][0] == x {
				found = true
			} else {
				found = false
			}

		}
		if !found {
			break
		}
	}
	for {
		y = rand.IntN(state.Bound_y-2) + 1
		found := false
		for i := range state.SnakePosition {
			if state.SnakePosition[i][1] == y {
				found = true
			} else {
				found = false
			}

		}
		if !found {
			break
		}
	}

	return x, y
}

func checkEating(state *model.Gamestate) {
	ele1 := state.SnakePosition[0]
	ele2 := state.SnakePosition[len(state.SnakePosition)-1]
	check := (state.Food_x == ele1[1] && state.Food_y == ele1[0]) || (state.Food_x == ele2[1] && state.Food_y == ele2[0])

	if check {
		state.SnakePosition = append(state.SnakePosition, []int{0, 0, 0})
		state.Food_x, state.Food_y = generateFood(state)
	}

	for i, val := range state.SnakePosition {
		if i == 0 {
			continue
		}
		if ele1[0] == val[0] && ele1[1] == val[1] {
			state.SnakePosition = [][]int{{3, 4, 1}, {3, 5, 1}}
			state.Game_over = true

		}
	}
}

func readInput(inputChan chan rune, s tcell.Screen) {
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

func updateDiraction(input rune, state *model.Gamestate) {
	switch input {
	case 's':
		state.Direction = 1
		state.SnakePosition[0][2] = 1 //down
	case 'w':
		state.Direction = 2
		state.SnakePosition[0][2] = 2 //up

	case 'd':
		state.Direction = 3
		state.SnakePosition[0][2] = 3 //right

	case 'a':
		state.Direction = 4
		state.SnakePosition[0][2] = 4 // left
	case 'r':
		state.Game_over = false
	}

}

func RunGame() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

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

	_, height := s.Size()

	state := initGameState(s)
	Food_x, Food_y := generateFood(&state)
	state.Food_x = Food_x
	state.Food_y = Food_y

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
	index := 1
	for {

		if index >= len(state.SnakePosition) {
			index = 0
		}
		s.Clear()
		select {
		case <-quitChan:
			return
		case input := <-inputChan:
			updateDiraction(input, &state)
		default:

			gameLoop(&state, s, tcell.StyleDefault, index)
			if err != nil {
				panic(err)
			}

			putString(s, 0, height-1, defStyle, "Score: "+strconv.Itoa(len(state.SnakePosition)-1))

			s.Show()

			time.Sleep(100 * time.Millisecond)

		}
		index++

	}

	// Clear screen

}

func putString(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for i, r := range str {
		s.SetContent(x+i, y, r, nil, style)
	}
}
