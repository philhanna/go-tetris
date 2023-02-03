package tetris

// TetrisType is an enumeration of the type/shape of a tetromino, not
// including orientation.
type TetrisType int8

const (
	TET_I TetrisType = iota
	TET_J
	TET_L
	TET_O
	TET_S
	TET_T
	TET_Z
)

// String returns a string representation of a tetris type
func (t TetrisType) String() string {
	switch t {
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
	return "?"
}
