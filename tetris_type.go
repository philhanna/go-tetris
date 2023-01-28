package tetris

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
func (tt TetrisType) TypeToCell() TetrisCell {
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
