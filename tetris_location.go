package tetris

import (
	"fmt"
)

// A TetrisLocation is a (row, column) pair. Negative numbers are
// allowed, because we need them for offsets.
type TetrisLocation struct {
	row int
	col int
}

// String returns the tetris location as a string {row, col}
func (tl TetrisLocation) String() string {
	return fmt.Sprintf("{%d, %d}", tl.row, tl.col)
}
