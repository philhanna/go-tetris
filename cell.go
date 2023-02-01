package tetris

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

// String returns a string representation of a TetrisCell
func (c TetrisCell) String() string {
	switch c {
	case TC_EMPTY:
		return " "
	case TC_CELLI:
		return "I"
	case TC_CELLJ:
		return "J"
	case TC_CELLL:
		return "L"
	case TC_CELLO:
		return "O"
	case TC_CELLS:
		return "S"
	case TC_CELLT:
		return "T"
	case TC_CELLZ:
		return "Z"
	default:
		return "?"
	}
}
