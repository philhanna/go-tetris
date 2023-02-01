package tetris

import (
	"math/rand"
)

// Tetrominio is an array of four tetris shapes, one for each of its
// orientations.
type Tetromino [NUM_ORIENTATIONS][NUM_CELLS]TetrisLocation

var tetrominos = map[TetrisType]Tetromino{
	TET_I: {
		{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		{{3, 0}, {3, 1}, {3, 2}, {3, 3}},
		{{0, 1}, {1, 1}, {2, 1}, {3, 1}}},
	TET_J: {
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 1}, {1, 1}, {2, 0}, {2, 1}}},
	TET_L: {
		{{0, 2}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 0}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 1}}},
	TET_O: {
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}}},
	TET_S: {
		{{0, 1}, {0, 2}, {1, 0}, {1, 1}},
		{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
		{{1, 1}, {1, 2}, {2, 0}, {2, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 1}}},
	TET_T: {
		{{0, 1}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {1, 2}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 1}, {1, 0}, {1, 1}, {2, 1}}},
	TET_Z: {
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		{{0, 2}, {1, 1}, {1, 2}, {2, 1}},
		{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
		{{0, 1}, {1, 0}, {1, 1}, {2, 0}}},
}

// RandomTetronimo returns a randomly chosen one of the tetronimos
func RandomTetronimo() Tetromino {
	i := rand.Intn(NUM_TETROMINOS)
	keys := []TetrisType{TET_I, TET_J, TET_L, TET_O, TET_S, TET_T, TET_Z}
	key := keys[i]
	return tetrominos[key]
}
