package tetris

import "math/rand"

// TetrisBlock is a struct that contains information about a tetromino.
// Specifically, what type it is, what orientation it has, and where it
// is.
type TetrisBlock struct {
	blockType   TetrisType
	orientation int
	location    TetrisLocation
}

func RandomBlock(cols int) *TetrisBlock {

	p := new(TetrisBlock)

	tetrisTypes := []TetrisType{TET_I, TET_J, TET_L, TET_O, TET_S, TET_T, TET_Z}
	i := rand.Intn(len(tetrisTypes))

	p.blockType = tetrisTypes[i]
	p.orientation = 0
	p.location.row = 0
	p.location.col = cols/2 - 2

	return p
}
