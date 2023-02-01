package tetris

import "fmt"

// A TetrisLocation is a row, column pair. Negative numbers allowed,
// because we need them for offsets.
type TetrisLocation struct {
	row int
	col int
}

// String returns a string representation of a TetrisLocation
func (loc TetrisLocation) String() string {
	return fmt.Sprintf("(%d,%d)", loc.row, loc.col)
}
