package tetris

import "math/rand"

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

// Fits determines whether this block can be placed on the board.  This
// means that the block's row and column are in bounds and that the
// board cell is not filled.
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

// RandomTetrominoType returns a random tetromino type
func RandomTetrominoType() TetrisType {
	r := rand.Intn(NUM_TETROMINOS)
	return TetrisType(r)
}

// NewFalling creates a new falling block and populate the next falling
// block with a random one.
func (tg *TetrisGame) NewFalling() {
	// Move next to falling
	tg.falling = tg.next

	// Update next to a new tetromino
	tg.next.typ = RandomTetrominoType()
	tg.next.ori = 0
	tg.next.loc.row = 0
	tg.next.loc.col = tg.cols/2 - 2
}

// DoGravityTick ticks gravity, and moves the block down if gravity
// should act.
func (tg *TetrisGame) DoGravityTick() {

	// Count down the clock
	tg.ticks_until_gravity--

	// If time is expired:
	if tg.ticks_until_gravity <= 0 {

		// Temporarily remove the falling block from its current
		// location and increment its row
		tg.Remove(tg.falling)
		tg.falling.loc.row++

		// If it still fits in the board, reset the clock according to
		// the gravity for this level
		if tg.Fits(tg.falling) {
			tg.ticks_until_gravity = GRAVITY_LEVEL[tg.level]
		} else {

			// Otherwise, restore the falling block's row, store it on
			// the pile, and create a new falling block (actually, just
			// move the "next" to "falling"
			tg.falling.loc.row--
			tg.Put(tg.falling)
			tg.NewFalling()
		}
		tg.Put(tg.falling)
	}
}

// Move moves the falling tetris block left (-1) or right (+1). If it
// doesn't fit, put it back.
func (tg *TetrisGame) Move(direction int) {

	// Temporarily remove the falling block and update its column
	// according to the desired direction.
	tg.Remove(tg.falling)
	tg.falling.loc.col += direction

	// If it doesn't fit in the new location, move it back
	if !tg.Fits(tg.falling) {
		tg.falling.loc.col -= direction
	}

	// Restore the falling block in the board
	tg.Put(tg.falling)
}

// Down sends the falling tetris block to the bottom.
func (tg *TetrisGame) Down() {

	// Remove the falling block, and move it down until it won't fit
	tg.Remove(tg.falling)
	for tg.Fits(tg.falling) {
		tg.falling.loc.row++
	}

	// Whoa, back up, we're as far down as we can go
	tg.falling.loc.row--

	// Put it back on the board
	tg.Put(tg.falling)
	tg.NewFalling()
}
