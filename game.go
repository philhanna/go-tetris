package tetris

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// TetrisGame is a class that holds all the information about the state
// of play
type TetrisGame struct {
	// Dimensions and contents of the board
	nRows int
	nCols int
	board [][]TetrisCell

	// Scoring information
	points int
	level  int

	// Falling block is the one currently going down.
	// Next block is the one that will be fallingBlock after this one.
	// Stored is the block that you can swap out.
	fallingBlock *TetrisBlock
	nextBlock    *TetrisBlock
	storedBlock  *TetrisBlock

	// Number of game ticks until the block will move down
	ticksRemaining int

	// Number of lines until you advance to the next level.
	linesRemaining int
}

// NewTetrisGame is the constructor for a TetrisGame object
func NewTetrisGame(nRows, nCols int) TetrisGame {
	pGame := new(TetrisGame)
	pGame.Init(nRows, nCols)
	return *pGame
}

// Adjust the score for the game, given how many lines were just cleared.
func (pGame *TetrisGame) AdjustScore(linesCleared int) {
	// Note: the array is limited to a maximum of four lines cleared,
	// which seems OK because that is as big as a tetromino can get
	lineMultiplier := []int{0, 40, 100, 300, 1200}

	// Add the points for the number of lines cleared to the game's
	// points variable
	multiplier := lineMultiplier[linesCleared]
	nextLevel := pGame.level + 1
	points := multiplier * nextLevel
	pGame.points += points

	// Update the lines remaining variable.
	if linesCleared >= pGame.linesRemaining {
		pGame.level = min(MAX_LEVEL, pGame.level+1)
		pGame.linesRemaining = LINES_PER_LEVEL - linesCleared
	} else {
		pGame.linesRemaining = linesCleared
	}
}

// Find rows that are filled, remove them, shift, and return the number
// of cleared rows.
func (pGame *TetrisGame) CheckLines() int {

	// Save and restore the falling block
	pGame.Remove(pGame.fallingBlock)

	// Find the number of lines that have been cleared in this method
	nLines := 0
	for i := pGame.nRows - 1; i >= 0; i-- {
		if pGame.LineFull(i) {
			pGame.ShiftLines(i)
			i++ // Do this line again because they've shifted
			nLines++
		}
	}

	// Restore the falling block
	pGame.Put(pGame.fallingBlock)

	return nLines
}

// DoGravityTick does a single game tick: process gravity, user input,
// and score.  Return true if the game is still running, false if it is
// over.
func (pGame *TetrisGame) DoGravityTick() {
	pGame.ticksRemaining--
	if pGame.ticksRemaining <= 0 {
		pGame.Remove(pGame.fallingBlock)
		pGame.fallingBlock.location.row++
		if pGame.Fits(pGame.fallingBlock) {
			pGame.ticksRemaining = GRAVITY_LEVEL[pGame.level]
		} else {
			pGame.fallingBlock.location.row--
			pGame.Put(pGame.fallingBlock)
		}
		pGame.Put(pGame.fallingBlock)
	}
}

// Send the falling tetris block to the bottom.
func (pGame *TetrisGame) Down() {
	pGame.Remove(pGame.fallingBlock)
	for pGame.Fits(pGame.fallingBlock) {
		pGame.fallingBlock.location.row++
	}
	pGame.fallingBlock.location.row--
	pGame.Put(pGame.fallingBlock)
	pGame.MakeNewBlocks()
}

// Checks whether a block can be placed on the board
func (pGame *TetrisGame) Fits(block *TetrisBlock) bool {
	for i := 0; i < NUM_CELLS; i++ {
		location := Tetrominos[block.blockType][block.orientation][i]
		row := block.location.row + location.row
		col := block.location.col + location.col

		// If the new location is outside the board, no fit.
		if ok, _ := pGame.WithinBounds(row, col); !ok {
			return false
		}

		// If the cell at the new location is not empty, no fit
		newCell := pGame.Get(row, col)
		if newCell != TC_EMPTY {
			return false
		}
	}
	return true
}

// Returns true if the game is over
func (pGame *TetrisGame) GameOver() bool {

	// Save and restore falling block at the end of this method.
	pGame.Remove(pGame.fallingBlock)
	over := false

	// If the top two rows are not clear, the game is over
	for i := 0; i < 2; i++ {
		for j := 0; j < pGame.nCols; j++ {
			cell := pGame.Get(i, j)
			if cell != TC_EMPTY {
				over = true
			}
		}
	}
	pGame.Put(pGame.fallingBlock)
	return over
}

// Get returns the cell at the given row and column.
func (pGame *TetrisGame) Get(row, col int) TetrisCell {
	if ok, _ := pGame.WithinBounds(row, col); !ok {
		return TC_EMPTY
	}
	return pGame.board[row][col]
}

// HandleMove performs the action specified by the move.
func (pGame *TetrisGame) HandleMove(move TetrisMove) {
	switch move {
	case TM_LEFT:
		pGame.Move(-1)
	case TM_RIGHT:
		pGame.Move(1)
	case TM_DROP:
		pGame.Down()
	case TM_CLOCK:
		pGame.Rotate(1)
	case TM_COUNTER:
		pGame.Rotate(-1)
	case TM_HOLD:
		pGame.Hold()
	}
}

// Swap the falling block with the block in the hold buffer.
func (pGame *TetrisGame) Hold() {
	pGame.Remove(pGame.fallingBlock)
	if pGame.storedBlock == nil {

		// There is no stored block. Just copy the falling one.

		pGame.storedBlock = pGame.fallingBlock
	} else {

		// There is a stored block. Swap with the falling one.

		typ := pGame.fallingBlock.blockType
		ori := pGame.fallingBlock.orientation

		pGame.fallingBlock.blockType = pGame.storedBlock.blockType
		pGame.fallingBlock.orientation = pGame.storedBlock.orientation

		pGame.storedBlock.blockType = typ
		pGame.storedBlock.orientation = ori
		for !pGame.Fits(pGame.fallingBlock) {
			pGame.fallingBlock.location.row--
		}
	}
	pGame.Put(pGame.fallingBlock)
}

// Init initializes the data in a TetrisGame object
func (pGame *TetrisGame) Init(nRows, nCols int) {

	// Initialize the boilerplate fields
	pGame.nRows = nRows
	pGame.nCols = nCols
	pGame.board = NewBoard(nRows, nCols)
	pGame.points = 0
	pGame.level = 0
	pGame.nextBlock = nil
	pGame.storedBlock = nil
	pGame.ticksRemaining = GRAVITY_LEVEL[0]
	pGame.linesRemaining = LINES_PER_LEVEL

	// Initialize the random number generator so that we can generate
	// the first two random tetrominos
	rand.Seed(time.Now().UnixNano())

	// Now run MakeNewBlocks twice to initialize the falling and stored
	// pointers
	pGame.MakeNewBlocks()
	pGame.MakeNewBlocks()
}

// Returns true if line i is full
func (pGame *TetrisGame) LineFull(i int) bool {
	for j := 0; j < pGame.nCols; j++ {
		cell := pGame.Get(i, j)
		if cell == TC_EMPTY {
			return false
		}
	}
	return true
}

// MakeNewBlocks moves the next block to the falling block and creates a
// new randomly chosen block
func (pGame *TetrisGame) MakeNewBlocks() {
	pGame.fallingBlock = pGame.nextBlock
	pGame.nextBlock = RandomBlock(pGame.nCols)
}

// Move moves the falling tetris block left (-1) or right (+1), given
// that it fits
func (pGame *TetrisGame) Move(direction int) {
	pGame.Remove(pGame.fallingBlock)
	pGame.fallingBlock.location.col += direction
	if !pGame.Fits(pGame.fallingBlock) {
		pGame.fallingBlock.location.col -= direction
	}
	pGame.Put(pGame.fallingBlock)
}

// NewBoard allocates memory for a rows*cols board.
func NewBoard(nRows, nCols int) [][]TetrisCell {

	// Define the board as a slice of slices of TetrisCells
	board := make([][]TetrisCell, nRows)

	// Outer loop creates a row at a time
	for row := range board {
		board[row] = make([]TetrisCell, nCols)
	}

	// Initialize all cells to empty
	for row := 0; row < nRows; row++ {
		for col := 0; col < nCols; col++ {
			board[row][col] = TC_EMPTY
		}
	}
	return board
}

// Put places a block onto the board
func (pGame *TetrisGame) Put(block *TetrisBlock) {
	for i := 0; i < NUM_CELLS; i++ {
		location := Tetrominos[block.blockType][block.orientation][i]
		newRow := block.location.row + location.row
		newCol := block.location.col + location.col
		pGame.Set(newRow, newCol, TC_EMPTY)
	}
}

// Clears a block off the board
func (pGame *TetrisGame) Remove(block *TetrisBlock) {
	for i := 0; i < NUM_CELLS; i++ {
		location := Tetrominos[block.blockType][block.orientation][i]
		newRow := block.location.row + location.row
		newCol := block.location.col + location.col
		pGame.Set(newRow, newCol, TypeToCell(block.blockType))
	}
}

// Rotate rotates the falling block in either direction (+/-1), given
// that it fits
func (pGame *TetrisGame) Rotate(direction int) {
	pGame.Remove(pGame.fallingBlock)
	for {
		pGame.fallingBlock.orientation =
			(pGame.fallingBlock.orientation + direction) % NUM_ORIENTATIONS

		// If the new orientation fits, we're done
		if pGame.Fits(pGame.fallingBlock) {
			break
		}

		// Otherwise, try moving it to the left to make it fit
		pGame.fallingBlock.location.col--
		if pGame.Fits(pGame.fallingBlock) {
			break
		}

		// Finally, try moving it to the right to make it fit
		pGame.fallingBlock.location.col += 2
		if pGame.Fits(pGame.fallingBlock) {
			break
		}

		// Put it back in its original location and try the next
		// orientation. Worst case, we come back to the original
		// orientation and it fits, so this loop will terminate
		pGame.fallingBlock.location.col--
	}
	pGame.Put(pGame.fallingBlock)
}

// Set sets the cell at the given row and column.
func (pGame *TetrisGame) Set(row, col int, value TetrisCell) error {
	if _, err := pGame.WithinBounds(row, col); err != nil {
		return err
	}
	pGame.board[row][col] = value
	return nil
}

// Shift every row above r down one
func (pGame *TetrisGame) ShiftLines(r int) {
	for i := r - 1; i >= 0; i-- {
		for j := 0; j < pGame.nCols; j++ {
			cell := pGame.Get(i, j)
			pGame.Set(i+1, j, cell)
			pGame.Set(i, j, TC_EMPTY)
		}
	}
}

// String returns a string representation of the game object
func (pGame *TetrisGame) String() string {

	var sb strings.Builder

	sb.WriteString("Game:\n")
	sb.WriteString(fmt.Sprintf("  nRows: %d\n", pGame.nRows))
	sb.WriteString(fmt.Sprintf("  nCols: %d\n", pGame.nCols))
	sb.WriteString("  board: {\n")

	// Write each row
	for _, cells := range pGame.board {
		sb.WriteString("    {")
		// Write each column in this row
		for _, cell := range cells {
			sb.WriteString(cell.String())
		}
		sb.WriteString("}\n")
	}
	sb.WriteString("  }\n")
	sb.WriteString(fmt.Sprintf("  points: %d\n", pGame.points))
	sb.WriteString(fmt.Sprintf("  level: %d\n", pGame.level))

	blockString := func(p *TetrisBlock) string {
		if p != nil {
			return p.String()
		}
		return "<nil>"
	}
	sb.WriteString(fmt.Sprintf("  falling: %s\n", blockString(pGame.fallingBlock)))
	sb.WriteString(fmt.Sprintf("  next:    %s\n", blockString(pGame.nextBlock)))
	sb.WriteString(fmt.Sprintf("  stored:  %s\n", blockString(pGame.storedBlock)))
	sb.WriteString(fmt.Sprintf("  ticksRemaining: %d\n", pGame.ticksRemaining))
	sb.WriteString(fmt.Sprintf("  linesRemaining: %d", pGame.linesRemaining))

	s := sb.String()
	return s
}

// Do a single game tick: process gravity, user input, and score.
// Return true if the game is still running, false if it is over.
func (pGame *TetrisGame) Tick(move TetrisMove) bool {

	// Handle gravity
	pGame.DoGravityTick()

	// Handle input
	pGame.HandleMove(move)

	// Check for cleared lines
	linesCleared := pGame.CheckLines()
	pGame.AdjustScore(linesCleared)

	// Return whether the game will continue (NOT whether it's over)
	gameOver := pGame.GameOver()

	return !gameOver

}

// Converts a TetrisType to a TetrisCell
func TypeToCell(typ TetrisType) TetrisCell {
	switch typ {
	case TET_I:
		return TC_CELLI
	case TET_J:
		return TC_CELLJ
	case TET_L:
		return TC_CELLL
	case TET_O:
		return TC_CELLO
	case TET_S:
		return TC_CELLS
	case TET_T:
		return TC_CELLT
	case TET_Z:
		return TC_CELLZ
	default:
		return TC_EMPTY
	}
}

// WithinBounds returns an error if the specified row or column is
// contained in the board
func (pGame *TetrisGame) WithinBounds(row, col int) (bool, error) {
	switch {
	case row < 0, row >= pGame.nRows:
		errmsg := fmt.Sprintf("row %d is not >= %d and < %d", row, 0, pGame.nRows)
		return false, errors.New(errmsg)
	case col < 0, col >= pGame.nCols:
		errmsg := fmt.Sprintf("col %d is not >= %d and < %d", col, 0, pGame.nCols)
		return false, errors.New(errmsg)
	default:
		return true, nil
	}
}
