package main

import (
	"log"
	"tetris"
	"time"

	gc "github.com/rthornton128/goncurses"
)

const COLS_PER_CELL = 2

// Main tetris game
func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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

	// Deinitialize curses when this function is done
	defer stdscr.Clear()
	defer gc.End()

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
	score, _ := gc.NewWindow(6, 10, 14, w+3)

	// Game loop
	for running {
		running = tg.Tick(move)
		DisplayBoard(board, tg)
		DisplayPiece(next, *tg.NextBlock)
		DisplayScore(score, tg)
		gc.Update()
		Sleep(10)

		ch := stdscr.GetChar()
		switch ch {
		case gc.KEY_LEFT:
			move = tetris.TM_LEFT
		case gc.KEY_RIGHT:
			move = tetris.TM_RIGHT
		case gc.KEY_UP:
			move = tetris.TM_CLOCK
		case gc.KEY_DOWN:
			move = tetris.TM_COUNTER
		case 'q':
			running = false
			move = tetris.TM_NONE
		case ' ':
			move = tetris.TM_DROP
		case 'p':
			doPause(board, tg, stdscr)
			move = tetris.TM_NONE
		case 'b':
			doBoss(board, tg, stdscr)
			move = tetris.TM_NONE
		default:
			move = tetris.TM_NONE
		}
	}

	// Display end of game message
	stdscr.Move(24, 0)
	stdscr.Printf("Game over!\n")
	stdscr.Printf("You finished with %d points on level %d.\n", tg.Points, tg.Level)
	stdscr.Timeout(-1)
	stdscr.GetChar()
}

// Handles a pause request
func doPause(win *gc.Window, tg *tetris.Game, stdscr *gc.Window) {
	win.Clear()
	win.Box(0, 0)
	y := tg.NRows / 2
	x := (tg.NCols*COLS_PER_CELL - 6) / 2
	win.Move(y, x)
	win.Print("PAUSED")
	win.Refresh()
	stdscr.Timeout(-1)
	stdscr.GetChar()
	stdscr.Timeout(0)
}

// Handles a boss request
func doBoss(win *gc.Window, tg *tetris.Game, stdscr *gc.Window) {
	stdscr.Clear()
	// insert boss here
	const bossImage = `
user@workstation-312:~/Documents/presentation $ ls -l
total 528
drwxr-xr-x 2 user users   4096 Jun  9 17:05 .
drwxr-xr-x 4 user users   4096 Jun 10 09:52 ..
-rw-r--r-- 1 user users  88583 Jun  9 14:13 figure1.png
-rw-r--r-- 1 user users  65357 Jun  9 15:40 figure2.png
-rw-r--r-- 1 user users   4469 Jun  9 16:17 presentation.aux
-rw-r--r-- 1 user users  42858 Jun  9 16:17 presentation.log
-rw-r--r-- 1 user users   2516 Jun  9 16:17 presentation.nav
-rw-r--r-- 1 user users    183 Jun  9 16:17 presentation.out
-rw-r--r-- 1 user users 349607 Jun  9 16:17 presentation.pdf
-rw-r--r-- 1 user users      0 Jun  9 16:17 presentation.snm
-rw-r--r-- 1 user users   9284 Jun  9 17:05 presentation.tex
-rw-r--r-- 1 user users    229 Jun  9 16:17 presentation.toc
user@workstation-312:~/Documents/presentation $
`
	stdscr.Print(bossImage)
	stdscr.Refresh()
	stdscr.Timeout(-1)
	stdscr.GetChar()
	stdscr.Timeout(0)
	stdscr.Clear()
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
// First, it draws a box around the border of the window.
// Then, in a nested loop of all columns within all rows:
// 1. At the beginning of each row, moves the cursor to the row below
// 2. For each column, gets the cell on the board at the row, col
// 
func DisplayBoard(win *gc.Window, tg *tetris.Game) {
	win.Box(0, 0)
	for i := 0; i < tg.NRows; i++ {
		win.Move(1+i, 1)
		for j := 0; j < tg.NCols; j++ {
			cell := tg.Get(i, j)
			if cell != tetris.TC_EMPTY {
				AddBlock(win, cell)
			} else {
				AddEmpty(win)
			}
		}
	}
	win.NoutRefresh()
}

// AddBlock draws a cell with the right color
func AddBlock(win *gc.Window, cell tetris.Cell) {
	var ach gc.Char
	for i := 0; i < COLS_PER_CELL; i++ {
		ach = ' ' | gc.A_REVERSE | gc.ColorPair(int16(cell))
		win.AddChar(ach)
	}
}

// AddEmpty erases a cell
func AddEmpty(win *gc.Window) {
	var ach gc.Char
	for i := 0; i < COLS_PER_CELL; i++ {
		ach = ' '
		win.AddChar(ach)
	}
}

// DisplayPiece displays a tetris piece in a dedicated window.
func DisplayPiece(win *gc.Window, block tetris.Block) {
	win.Clear()
	win.Box(0, 0)
	if block.BlockType == -1 {
		win.NoutRefresh()
		return
	}
	for b := 0; b < tetris.NUM_CELLS; b++ {
		location := tetris.Tetrominos[block.BlockType][block.Orientation][b]
		y := location.Row + 1
		x := location.Col*COLS_PER_CELL + 1
		win.Move(y, x)
		cell := tetris.TypeToCell(block.BlockType)
		AddBlock(win, cell)
	}
	win.NoutRefresh()
}

// DisplayScore displays score information in a dedicated window.
func DisplayScore(win *gc.Window, tg *tetris.Game) {
	win.Clear()
	win.Box(0, 0)
	win.Printf("Score\n%d\n", tg.Points)
	win.Printf("Level\n%d\n", tg.Level)
	win.Printf("Lines\n%d\n", tg.LinesRemaining)
	win.NoutRefresh()
}

// Sleep sleeps for the specified number of milliseconds
func Sleep(millis int) {
	n := time.Duration(millis) * time.Millisecond
	time.Sleep(n)
}
