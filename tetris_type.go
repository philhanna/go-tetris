package tetris

// TetrisType is a type/shape of a tetromino, not including orientation.
type TetrisType int

const (
	TET_I = iota
	TET_J
	TET_L
	TET_O
	TET_S
	TET_T
	TET_Z
)

// ToCell converts a tetromino type to its corresponding cell.
func (tt TetrisType) ToCell() TetrisCell {
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
	}
	return 0
}

// String returns the representation of this object as a string.
func (tt TetrisType) String() string {
	switch tt {
	case TET_I:
		return "I"
	case TET_J:
		return "J"
	case TET_L:
		return "L"
	case TET_O:
		return "O"
	case TET_S:
		return "S"
	case TET_T:
		return "T"
	case TET_Z:
		return "Z"
	}
	return ""
}
