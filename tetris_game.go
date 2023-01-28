package tetris

// TetrisGame is a game object
type TetrisGame struct {
	rows  int
	cols  int
	board []TetrisBlock

	// Scoring information
	points int
	level  int

	// Falling block is the one currently going down.  Next block is the
	// one that will be falling after this one.  Stored is the block
	// that you can swap out.
	falling TetrisBlock
	next    TetrisBlock
	stored  TetrisBlock

	// Number of game ticks until the block will move down
	ticksTillGravity int

	// Number of lines until you advance to the next level
	linesRemaining int
}

// Get returns the block at the given row and col.
func (tg *TetrisGame) Get(row int, col int) TetrisBlock {
	return tg.board[tg.cols*row+col]
}

// Set sets the block at the given row and col
func (tg *TetrisGame) Set(row int, col int, value TetrisBlock) {
	tg.board[tg.cols*row+col] = value
}

// Check checks whether a row and col are in bounds.
func (tg *TetrisGame) Check(row, col int) bool {
	return 0 <= row && row < tg.rows && 0 <= col && col < tg.cols
}

// Put places a block onto the board.
func (tg *TetrisGame) Put(block TetrisBlock) {
	for i := 0; i < TETRIS; i++ {
		cell := TETROMINOS[block.typ][block.orientation][i]
		row := block.location.row + cell.row
		col := block.location.col + cell.col
		value := block
		tg.Set(row, col, value)
	}
}

// Remove clears a block out of the board
func (tg *TetrisGame) Remove(block TetrisBlock) {
	for i := 0; i < TETRIS; i++ {
		cell := TETROMINOS[block.typ][block.orientation][i]
		row := block.location.row + cell.row
		col := block.location.col + cell.col
		tg.Set(row, col, TC_EMPTY)
	}
}

// Fits checks whether a block can be placed on the board.
func (tg *TetrisGame) Fits(block TetrisBlock) bool {
	for i := 0; i < TETRIS; i++ {
		cell := TETROMINOS[block.typ][block.orientation][i]
		row := block.location.row + cell.row
		col := block.location.col + cell.col
		if !tg.Check(row, col) || block.typ.TypeToCell(tg.Get(row, col)) {
			return false
		}
	}
	return true
}
