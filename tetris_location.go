package tetris

// A TetrisLocation is a row, column pair. Negative numbers allowed,
// because we need them for offsets.
type TetrisLocation struct {
	row int
	col int
}
