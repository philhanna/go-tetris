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

// Check determines whether a row and column are in bounds.
func (tg *TetrisGame) Check(row int, col int) bool {
	rowInBounds := 0 <= row && row < tg.rows
	colInBounds := 0 <= col && col < tg.cols
	return rowInBounds && colInBounds
}

// Put places a block on the board
func (tg *TetrisGame) Put(block TetrisBlock) {
	for i := 0; i < TETRIS; i++ {
		cell := TETROMINOS[block.typ][block.ori][i]
		row := block.loc.row + cell.row
		col := block.loc.col + cell.col
		tg.Set(row, col, block.typ.ToCell())
	}
}

// Remove clears a block from the board
func (tg *TetrisGame) Remove(block TetrisBlock) {
	for i := 0; i < TETRIS; i++ {
		cell := TETROMINOS[block.typ][block.ori][i]
		row := block.loc.row + cell.row
		col := block.loc.col + cell.col
		tg.Set(row, col, TC_EMPTY)
	}
}

// Fits determines whether this block can be placed on the board
// (End of game if so?)
func (tg *TetrisGame) Fits(block TetrisBlock) bool {
	for i := 0; i < TETRIS; i++ {
		cell := TETROMINOS[block.typ][block.ori][i]
		row := block.loc.row + cell.row
		col := block.loc.col + cell.col
		inBounds := tg.Check(row, col)
		boardCell := tg.Get(row, col)
		isFilled := boardCell.IsFilled()
		if !inBounds || isFilled {
			return false
		}
	}
	return true
}
