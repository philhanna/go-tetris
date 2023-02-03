package tetris

import "fmt"

// A Location is a row, column pair. Negative numbers allowed,
// because we need them for offsets.
type Location struct {
	row int
	col int
}

// String returns a string representation of a Location
func (loc Location) String() string {
	return fmt.Sprintf("(%d,%d)", loc.row, loc.col)
}
