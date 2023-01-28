package tetris

import (
	"fmt"
	"math/rand"
	"time"
)

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
	linesRemaining int
}

func (tg *TetrisGame) Init(rows int, cols int) {
	tg.rows = rows
	tg.cols = cols
	tg.board = make([]TetrisCell, (rows * cols))
	for i := 0; i < rows*cols; i++ {
		tg.board[i] = TC_EMPTY
	}
	tg.points = 0
	tg.level = 0
	tg.ticks_until_gravity = GRAVITY_LEVEL[0]
	tg.linesRemaining = LINES_PER_LEVEL
	rand.Seed(time.Now().UnixNano())
	tg.NewFalling()
	tg.NewFalling()
	tg.stored.typ = -1
	tg.stored.ori = 0
	tg.stored.loc.row = 0
	tg.next.loc.col = tg.cols/2 - 2
	fmt.Printf("%d", tg.falling.loc.col)
}

// Create is the TetrisGame constructor
func Create(rows int, cols int) *TetrisGame {
	tg := new(TetrisGame)
	tg.Init(rows, cols)
	return tg
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
			// Move the "next" to "falling"
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

// Rotate rotates the falling block in either direction (+/-1).
func (tg *TetrisGame) Rotate(direction int) {

	tg.Remove(tg.falling)

	for {
		tg.falling.ori = (tg.falling.ori + direction) % NUM_ORIENTATIONS

		// If the new orientation fits, we're done.
		if tg.Fits(tg.falling) {
			break
		}

		// Otherwise, try moving left to make it fit
		tg.falling.loc.col--
		if tg.Fits(tg.falling) {
			break
		}

		// Finally, try moving right to make it fit
		tg.falling.loc.col += 2
		if tg.Fits(tg.falling) {
			break
		}

		// Put it back in its original location and try the next
		// orientation. Worst case, we come back to the original
		// orientation and it fits, so this loop will terminate.
		tg.falling.loc.col--
	}

	tg.Put(tg.falling)
}

// Hold swaps the falling block with the block in the hold buffer.
// TODO check on type == -1
func (tg *TetrisGame) Hold() {

	tg.Remove(tg.falling)
	if tg.stored.typ == -1 {
		tg.stored = tg.falling
		tg.NewFalling()
	} else {
		typ := tg.falling.typ
		ori := tg.falling.ori
		tg.falling.typ = tg.stored.typ
		tg.falling.ori = tg.stored.ori
		tg.stored.typ = typ
		tg.stored.ori = ori
		for !tg.Fits(tg.falling) {
			tg.falling.loc.row--
		}
	}
	tg.Put(tg.falling)
}

// HandleMove performs the action specified by the move
func (tg *TetrisGame) HandleMove(move TetrisMove) {
	switch move {
	case TM_LEFT:
		tg.Move(-1)
	case TM_RIGHT:
		tg.Move(1)
	case TM_DROP:
		tg.Down()
	case TM_COUNTER:
		tg.Rotate(-1)
	case TM_HOLD:
		tg.Hold()
	}
}

// LineFull returns true if line i is full
func (tg *TetrisGame) LineFull(i int) bool {
	for j := 0; j < tg.cols; j++ {
		ijCell := tg.Get(i, j)
		if ijCell.IsEmpty() {
			return false
		}
	}
	return true
}

// ShiftLines shift every row above r down one.
func (tg *TetrisGame) ShiftLines(r int) {
	for i := r - 1; i >= 0; i-- {
		for j := 0; j < tg.cols; j++ {
			tg.Set(i+1, j, tg.Get(i, j))
			tg.Set(i, j, TC_EMPTY)
		}
	}
}

// CheckLines find rows that are filled, removes them, shift, and
// returns the number of cleared rows.
func (tg *TetrisGame) CheckLines() int {
	nLines := 0
	tg.Remove(tg.falling) // Don't want to mess up falling block

	for i := tg.rows - 1; i >= 0; i-- {
		if tg.LineFull(i) {
			tg.ShiftLines(i)
			i++ // Do this line over again since they're shifted
			nLines++
		}
	}

	tg.Put(tg.falling)
	return nLines
}

// AdjustScore adjusts the score for the game, given how many lines
// were just cleared
func (tg *TetrisGame) AdjustScore(linesCleared int) {
	lineMultipliers := []int{0, 40, 100, 300, 1200}
	lineMultiplier := lineMultipliers[linesCleared]
	nextLevel := tg.level + 1
	addedPoints := lineMultiplier * nextLevel
	tg.points += addedPoints
	if linesCleared >= tg.linesRemaining {
		tg.level = MIN(MAX_LEVEL, tg.level+1)
		linesCleared -= tg.linesRemaining
		tg.linesRemaining = LINES_PER_LEVEL - linesCleared
	} else {
		tg.linesRemaining -= linesCleared
	}
}
