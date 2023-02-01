package tetris

// ---------------------------------------------------------------------
//                       Enumerations
// ---------------------------------------------------------------------

// TetrisType is an enumeration of the type/shape of a tetromino, not
// including orientation.
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

// TetrisCell is an enumeration of 1x1 blocks within a tetris board.
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

// TetrisMove is an enumeration of all possible moves to give as input
// to the game
type TetrisMove int

const (
	TM_LEFT TetrisMove = iota
	TM_RIGHT
	TM_CLOCK
	TM_COUNTER
	TM_DROP
	TM_HOLD
	TM_NONE
)

// ---------------------------------------------------------------------
//                       Data structures
// ---------------------------------------------------------------------

// A TetrisLocation is a row, column pair. Negative numbers allowed,
// because we need them for offsets.
type TetrisLocation struct {
	row int
	col int
}

// TetrisBlock is a struct that contains information about a tetromino.
// Specifically, what type it is, what orientation it has, and where it
// is.
type TetrisBlock struct {
	blockType   TetrisType
	orientation int
	location    TetrisLocation
}
