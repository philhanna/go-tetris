package tetris

// ---------------------------------------------------------------------
//                       Enumerations
// ---------------------------------------------------------------------

// TetrisType is an enumeration of the type/shape of a tetromino, not
// including orientation.
type TetrisType uint8

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
type TetrisCell uint8

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
type TetrisMove uint8

const (
	TM_LEFT TetrisMove = iota
	TM_RIGHT
	TM_CLOCK
	TM_COUNTER
	TM_DROP
	TM_HOLD
	TM_NONE
)