package main

import (
	"fmt"
	"os"
	"tetris/tetris"
	// "tetris/tetris"
	// ncurses "github.com/rthornton128/goncurses"
)

const (
	// 2 columns per cell makes the game much nicer.
	COLS_PER_CELL = 2
)

// Main tetris game
func main() {

	var (
		tg   tetris.Game
		argc = len(os.Args)
	)

	// Load file if given a filename.
	if argc >= 2 {
		// NOP for now
	} else {
		tg = tetris.NewGame(22, 10)
	}

	/*
		stdscr, err := ncurses.Init()
		if err != nil {
			log.Fatal(err)
		}
		defer ncurses.End()
	*/

	// Dummy usage of unused variables
	if true {
		fmt.Println("tg:", tg)
		// fmt.Println("stdscr:", stdscr)
	}
}
