package tetris

import (
	"fmt"
)

// A TetrisBlock is a struct that contains information about a
// tetromino.  Specifically, what type it is, what orientation it has,
// and where it is.

type TetrisBlock struct {
	typ TetrisType
	ori int
	loc TetrisLocation
}

// NewTetrisBlock is a constructor for a TetrisBlock
func NewTetrisBlock(typ TetrisType, ori int, row int, col int) TetrisBlock {
	tb := new(TetrisBlock)
	tb.typ = typ
	tb.ori = ori
	tb.loc.row = row
	tb.loc.col = col
	return *tb
}

// String returns a string representation of the block
func (tb TetrisBlock) String() string {
	s := fmt.Sprintf("TetrisBlock{type=%v, ori=%d, location=%v}", tb.typ, tb.ori, tb.loc.String())
	return s
}
