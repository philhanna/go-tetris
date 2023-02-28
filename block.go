package tetris

import (
	"fmt"
	"math/rand"
)

// Block is a struct that contains information about a tetromino.
// Specifically, what type it is, what orientation it has, and where it
// is.
type Block struct {
	BlockType   Type
	Orientation int
	Location    Location
}

// String returns a string representation of a Block
func (pBlock *Block) String() string {
	blockTypeStr := pBlock.BlockType.String()
	orientation := pBlock.Orientation
	locationStr := pBlock.Location.String()
	s := fmt.Sprintf("(%q,%d,%s", blockTypeStr, orientation, locationStr)
	return s
}

// Returns allocates memory for a new Block chosen at random and
// returns a pointer to it.
func RandomBlock(cols int) *Block {

	tetrisTypes := []Type{TET_I, TET_J, TET_L, TET_O, TET_S, TET_T, TET_Z}
	i := rand.Intn(len(tetrisTypes))
	location := Location{0, cols/2 - 2}
	block := Block{tetrisTypes[i], 0, location}

	return &block
}
