package main

import (
	"fmt"
	"log"

	// "os"
	"tetris/tetris"

	gc "github.com/rthornton128/goncurses"
)

const (
	COLS_PER_CELL = 2 // 2 columns per cell makes the game much nicer.
)

// Cell colors (used in initColors())
const (
	COLOR_BLACK int16 = iota
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA
	COLOR_CYAN
	COLOR_WHITE
)

// Main tetris game
func main() {

	// NCURSES initialization
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()
	defer stdscr.Clear()

	gc.CBreak(true)     // pass key presses to program, but not signals
	gc.Echo(false)      // don't echo key presses to screen
	stdscr.Keypad(true) // allow arrow keys
	stdscr.Timeout(0)   // no blocking on getch()
	gc.Cursor(0)        // cursor invisible
	initColors()        // setup tetris colors

	// Now we can create the game
	tg := tetris.NewGame(22, 10)

	// Create windows for each section of the interface.
	h := tg.NRows + 2
	w := COLS_PER_CELL * (tg.NCols + 1)
	board, _ := gc.NewWindow(h, w, 0, 0)
	next, _ := gc.NewWindow(6, 10, 0, w+3)
	hold, _ := gc.NewWindow(6, 10, 7, w+3)
	score, _ := gc.NewWindow(6, 10, 14, w+3)

	// TODO: Remove this section after all variables are referenced.
	if false {
		fmt.Println(tg)
		fmt.Println(board, next, hold, score)
	}
}

// Do the NCURSES initialization steps for color blocks.
func initColors() {
	gc.StartColor()
	gc.InitPair(int16(tetris.TC_CELLI), COLOR_CYAN, COLOR_BLACK)
	gc.InitPair(int16(tetris.TC_CELLJ), COLOR_BLUE, COLOR_BLACK)
	gc.InitPair(int16(tetris.TC_CELLL), COLOR_WHITE, COLOR_BLACK)
	gc.InitPair(int16(tetris.TC_CELLO), COLOR_YELLOW, COLOR_BLACK)
	gc.InitPair(int16(tetris.TC_CELLS), COLOR_GREEN, COLOR_BLACK)
	gc.InitPair(int16(tetris.TC_CELLT), COLOR_MAGENTA, COLOR_BLACK)
	gc.InitPair(int16(tetris.TC_CELLZ), COLOR_RED, COLOR_BLACK)
}
