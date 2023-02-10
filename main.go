package main

import (
	"fmt"
	"log"

	// "os"
	"tetris/tetris"

	gc "github.com/rthornton128/goncurses"
)

const COLS_PER_CELL = 2

// Main tetris game
func main() {
	
	var (
		tg      *tetris.Game
		running = true
		move    = tetris.TM_NONE
	)

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
	InitColors()        // setup tetris colors

	// Now we can create the game
	newGame := tetris.NewGame(22, 10)
	tg = &newGame

	// Create windows for each section of the interface.
	h := tg.NRows + 2
	w := COLS_PER_CELL * (tg.NCols + 1)
	board, _ := gc.NewWindow(h, w, 0, 0)
	next, _ := gc.NewWindow(6, 10, 0, w+3)
	hold, _ := gc.NewWindow(6, 10, 7, w+3)
	score, _ := gc.NewWindow(6, 10, 14, w+3)

	// Game loop
	for running {
		running = tg.Tick(move)
		DisplayBoard(board, tg)
		DisplayPiece(next, *tg.NextBlock)
		DisplayPiece(hold, *tg.StoredBlock)
		// displayScore(score, tg)
	}

	// TODO: Remove this section after all variables are referenced.
	if false {
		fmt.Println(next)
		fmt.Println(hold)
		fmt.Println(score)
		fmt.Println(move)
	}
}

// InitColors does the NCURSES initialization steps for color blocks.
func InitColors() {
	gc.StartColor()
	gc.InitPair(int16(tetris.TC_CELLI), gc.C_CYAN, gc.C_BLACK)
	gc.InitPair(int16(tetris.TC_CELLJ), gc.C_BLUE, gc.C_BLACK)
	gc.InitPair(int16(tetris.TC_CELLL), gc.C_WHITE, gc.C_BLACK)
	gc.InitPair(int16(tetris.TC_CELLO), gc.C_YELLOW, gc.C_BLACK)
	gc.InitPair(int16(tetris.TC_CELLS), gc.C_GREEN, gc.C_BLACK)
	gc.InitPair(int16(tetris.TC_CELLT), gc.C_MAGENTA, gc.C_BLACK)
	gc.InitPair(int16(tetris.TC_CELLZ), gc.C_RED, gc.C_BLACK)
}

// DisplayBoard prints the tetris board onto the ncurses window.
func DisplayBoard(w *gc.Window, tg *tetris.Game) {
	w.Box(0, 0)
	for i := 0; i < tg.NRows; i++ {
		w.Move(1+i, 1)
		for j := 0; j < tg.NCols; j++ {
			cell := tg.Get(i, j)
			if cell == tetris.TC_EMPTY {
				AddBlock(w, cell)
			} else {
				AddEmpty(w)
			}
		}
	}
	w.NoutRefresh()
}

// AddBlock draws a cell with the right color
func AddBlock(w *gc.Window, cell tetris.Cell) {
	var ach gc.Char
	for i := 0; i < COLS_PER_CELL; i++ {
		ach = ' ' | gc.A_REVERSE | gc.ColorPair(int16(cell))
		w.AddChar(ach)
	}
}

// AddEmpty erases a cell
func AddEmpty(w *gc.Window) {
	var ach gc.Char
	for i := 0; i < COLS_PER_CELL; i++ {
		ach = ' '
		w.AddChar(ach)
	}
}

// DisplayPiece displays a tetris piece in a dedicated window.
func DisplayPiece(w *gc.Window, block tetris.Block) {
	w.Clear()
	w.Box(0, 0)
	if block.BlockType == -1 {
		w.NoutRefresh()
		return
	}
	for b := 0; b < tetris.NUM_CELLS; b++ {
		location := tetris.Tetrominos[block.BlockType][block.Orientation][b]
		y := location.Row + 1
		x := location.Col * COLS_PER_CELL + 1
		w.Move(y, x)
		cell := tetris.TypeToCell(block.BlockType)	
		AddBlock(w, cell)
	}
	w.NoutRefresh()	
}