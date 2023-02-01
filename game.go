package tetris

// TetrisGame is a structure that holds all the information about the
// state of play
type TetrisGame struct {
	// Dimensions and contents of the board
	rows  int
	cols  int
	board [][]TetrisCell
}
