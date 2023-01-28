package tetris

// A TetrisLocation is a (row, column) pair. Negative numbers are
// allowed, because we need them for offsets.
type TetrisLocation struct {
	row int
	col int
}
