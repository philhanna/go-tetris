package tetris

// Type is an enumeration of the type/shape of a tetromino, not
// including orientation.
type Type int8

const (
	TET_I Type = iota
	TET_J
	TET_L
	TET_O
	TET_S
	TET_T
	TET_Z
)

// String returns a string representation of a tetris type
func (t Type) String() string {
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

// Returns the Cell corresponding to a Type
func TypeToCell(typ Type) Cell {
	switch typ {
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
