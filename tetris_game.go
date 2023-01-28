package tetris

// A TetrisGame is the game object
type TetrisGame struct {
	// Game board
	rows  int
	cols  int
	board []TetrisCell

	// Scoring information
	points int
	level  int

	// Falling block is the one currently going down,
	// Next block is the one that will be falling next.
	// Stored is the block you can swap out.
	falling TetrisBlock
	next    TetrisBlock
	stored  TetrisBlock

	// Number of game ticks until the block will move down
	ticks_until_gravity int

	// Number of lines until you advance to the next level.
	lines_remaining int
}

// Get returns the cell at the given row and column.
func (tg *TetrisGame) Get(row int, col int) TetrisCell {
	return tg.board[tg.cols*row+col]
}

// Set sets the cell at the given row and column.
func (tg *TetrisGame) Set(row int, col int, value TetrisCell) {
	index := tg.cols*row + col
	tg.board[index] = value
}
