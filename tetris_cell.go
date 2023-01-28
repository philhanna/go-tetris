package tetris

// TetrisCell is a 1x1 block within a tetris board.
type TetrisCell byte

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

func (tc TetrisCell) String() string {
	switch tc {
	case TC_EMPTY:
		return "-"
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
