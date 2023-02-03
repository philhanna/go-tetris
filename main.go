package main

import (
	"fmt"
	"tetris/tetris"
	// ncurses "github.com/rthornton128/goncurses"
)

const (
	// 2 columns per cell makes the game much nicer.
	COLS_PER_CELL = 2
)

// Main tetris game
func main() {
	var tg tetris.TetrisGame

	// Dummy usage of unused variables
	fmt.Println(tg)
}
