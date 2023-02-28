package tetris

import "fmt"

// A Location is a row, column pair. Negative numbers allowed,
// because we need them for offsets.
type Location struct {
	Row int
	Col int
}

// String returns a string representation of a Location
func (loc Location) String() string {
	return fmt.Sprintf("(%d,%d)", loc.Row, loc.Col)
}
