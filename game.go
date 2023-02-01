package tetris

// TetrisGame is a class that holds all the information about the
// state of play
type TetrisGame struct {
	// Dimensions and contents of the board
	rows  int
	cols  int
	board [][]TetrisCell

	// Scoring information
	points int
	level  int

	// Falling block is the one currently going down.
	// Next block is the one that will be falling after this one.
	// Stored is the block that you can swap out.
	falling *TetrisBlock
	next    *TetrisBlock
	stored  *TetrisBlock

	// Number of game ticks until the block will move down
	ticksUntilGravity int

	// Number of lines until you advance to the next level.
	linesRemaining int
}

// Get returns the cell at the given row and column.
func (g TetrisGame) Get(row, col int) TetrisCell {
	return g.board[row][col]
}
