package tetris

import (
	"fmt"
	"math/rand"
)

// TetrisBlock is a struct that contains information about a tetromino.
// Specifically, what type it is, what orientation it has, and where it
// is.
type TetrisBlock struct {
	blockType   TetrisType
	orientation int
	location    TetrisLocation
}

// String returns a string representation of a TetrisBlock
func (pBlock *TetrisBlock) String() string {
	blockTypeStr := pBlock.blockType.String()
	orientation := pBlock.orientation
	locationStr := pBlock.location.String()
	s := fmt.Sprintf("(%q,%d,%s", blockTypeStr, orientation, locationStr)
	return s
}

// Returns allocates memory for a new TetrisBlock chosen at random and
// returns a pointer to it.
func RandomBlock(cols int) *TetrisBlock {

	tetrisTypes := []TetrisType{TET_I, TET_J, TET_L, TET_O, TET_S, TET_T, TET_Z}
	i := rand.Intn(len(tetrisTypes))
	location := TetrisLocation{0, cols/2 - 2}
	block := TetrisBlock{tetrisTypes[i], 0, location}

	return &block
}
