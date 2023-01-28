package main

import "fmt"

const (
	TC_EMPTY_STR     = " "
	TC_BLOCK_STR     = "\u2588"
	TETRIS           = 4
	NUM_TETROMINOS   = 7
	NUM_ORIENTATIONS = 4
	MAX_LEVEL        = 19
	LINES_PER_LEVEL  = 10
)

// TetrisCell is a 1x1 block within a tetris board.
type TetrisCell int

const (
	TC_EMPTY TetrisCell = iota
	TC_CELLI
	TC_CELLJ
	TC_CELLL
	TC_CELLO
	TC_CELLS
	TC_CELLT
	TC_CELLZ
)

// TetrisType is a type/shape of a tetromino, not including orientation.
type TetrisType int

const (
	TET_I TetrisType = iota
	TET_J
	TET_L
	TET_O
	TET_S
	TET_T
	TET_Z
)

// TypeToCell converts a tetromino type to its corresponding cell.
func TypeToCell(tt TetrisType) TetrisCell {
	switch tt {
	case TET_I:
		return TC_CELLI
	case TET_J:
		return TC_CELLJ
	case TET_L:
		return TC_CELLL
	case TET_O:
		return TC_CELLO
	case TET_S:
		return TC_CELLS
	case TET_T:
		return TC_CELLT
	case TET_Z:
		return TC_CELLZ
	default:
		return TC_EMPTY
	}
}

// TetrisLocation is a row,column pair.  Negative numbers allowed,
// because we need them for offsets.
type TetrisLocation struct {
	row int
	col int
}

// TetrisBlock is a struct that contains information about a tetromino.
// Specifically, what type it is, what orientation it has, and where it
// is.
type TetrisBlock struct {
	typ int
	ori int
	loc TetrisLocation
}

// TetrisMove is all possible moves to give as input to the game.
type TetrisMove int

const (
	TM_LEFT = iota
	TM_RIGHT
	TM_CLOCK
	TM_COUNTER
	TM_DROP
	TM_HOLD
	TM_NONE
)

// TetrisGame is a game object
type TetrisGame struct {
	rows, cols int
	board      string

	// Scoring information
	points, level int

	// Falling block is the one currently going down.  Next block is the
	// one that will be falling after this one.  Stored is the block
	// that you can swap out.
	falling, next, stored TetrisBlock

	// Number of game ticks until the block will move down
	ticksTillGravity int

	// Number of lines until you advance to the next level
	linesRemaining int
}

// This array stores all necessary information about the cells that are
// filled by each tetromino.  The first index is the type of the
// tetromino (i.e. shape, e.g. I, J, Z, etc.).  The next index is the
// orientation (0-3).  The final array contains 4 tetris_location
// objects, each mapping to an offset from a point on the upper left
// that is the tetromino "origin".
var TETRONIMOS = [][][]TetrisLocation{
	// I block
	{
		{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		{{3, 0}, {3, 1}, {3, 2}, {3, 3}},
		{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
	},
	// J block
	{
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 1}, {1, 1}, {2, 0}, {2, 1}},
	},
	// L block
	{
		{{0, 2}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 0}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
	},
	// O block
	{
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
	},
	// S block
	{
		{{0, 1}, {0, 2}, {1, 0}, {1, 1}},
		{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
		{{1, 1}, {1, 2}, {2, 0}, {2, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
	},
	// T block
	{
		{{0, 1}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {1, 2}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 1}, {1, 0}, {1, 1}, {2, 1}},
	},
	// Z block
	{
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		{{0, 2}, {1, 1}, {1, 2}, {2, 1}},
		{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
		{{0, 1}, {1, 0}, {1, 1}, {2, 0}},
	},
}

// This array tells you how many ticks per gravity by level.  Decreases as level
// increases, to add difficulty.
var GRAVITY_LEVEL = []int{
	50, 48, 46, 44, 42, 40, 38, 36, 34, 32,
	30, 28, 26, 24, 22, 20, 16, 12, 8, 4,
}

func TYPE_TO_CELL(x int) int {
	return x + 1
}

func TC_IS_EMPTY(x string) bool {
	return x == TC_EMPTY_STR
}

func TC_IS_FILLED(x string) bool {
	return !TC_IS_EMPTY(x)
}

func MAX(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MIN(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(TETRONIMOS)
	fmt.Println(GRAVITY_LEVEL)
}
