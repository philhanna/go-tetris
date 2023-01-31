package tetris

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
