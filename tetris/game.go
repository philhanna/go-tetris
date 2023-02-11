package tetris

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Game is a structure in the model that holds all the information
// about the state of play.
type Game struct {
	// Dimensions and contents of the board
	NRows int
	NCols int
	Board [][]Cell

	// Scoring information
	Points int
	Level  int

	// Falling block is the one currently going down.
	// Next block is the one that will be FallingBlock after this one.
	// Stored is the block that you can swap out.
	FallingBlock *Block
	NextBlock    *Block
	StoredBlock  *Block

	// Number of game ticks until the block will move down
	TicksRemaining int

	// Number of lines until you advance to the next level.
	LinesRemaining int
}

// NewGame is the constructor for a Game object
func NewGame(nRows, nCols int) Game {
	pGame := new(Game)
	pGame.Init(nRows, nCols)
	return *pGame
}

// Adjust the score for the game, given how many lines were just cleared.
func (pGame *Game) AdjustScore(linesCleared int) {
	// Note: the array is limited to a maximum of four lines cleared,
	// which seems OK because that is as big as a tetromino can get
	lineMultiplier := []int{0, 40, 100, 300, 1200}

	// Add the points for the number of lines cleared to the game's
	// points variable
	multiplier := lineMultiplier[linesCleared]
	nextLevel := pGame.Level + 1
	points := multiplier * nextLevel
	pGame.Points += points

	// Update the lines remaining variable.
	if linesCleared >= pGame.LinesRemaining {
		pGame.Level = min(MAX_LEVEL, pGame.Level+1)
		pGame.LinesRemaining = LINES_PER_LEVEL - linesCleared
	} else {
		pGame.LinesRemaining = linesCleared
	}
}

// Find rows that are filled, remove them, shift, and return the number
// of cleared rows.
func (pGame *Game) CheckLines() int {

	// Save and restore the falling block
	pGame.Remove(pGame.FallingBlock)

	// Find the number of lines that have been cleared in this method
	nLines := 0
	for i := pGame.NRows - 1; i >= 0; i-- {
		if pGame.LineFull(i) {
			pGame.ShiftLines(i)
			i++ // Do this line again because they've shifted
			nLines++
		}
	}

	// Restore the falling block
	pGame.Put(pGame.FallingBlock)

	return nLines
}

// DoGravityTick does a single game tick: process gravity, user input,
// and score.  Return true if the game is still running, false if it is
// over.
func (pGame *Game) DoGravityTick() {
	pGame.TicksRemaining--
	if pGame.TicksRemaining <= 0 {
		pGame.Remove(pGame.FallingBlock)
		pGame.FallingBlock.Location.Row++
		if pGame.Fits(pGame.FallingBlock) {
			pGame.TicksRemaining = GRAVITY_LEVEL[pGame.Level]
		} else {
			pGame.FallingBlock.Location.Row--
			pGame.Put(pGame.FallingBlock)
		}
		pGame.Put(pGame.FallingBlock)
	}
}

// Send the falling tetris block to the bottom.
func (pGame *Game) Down() {
	pGame.Remove(pGame.FallingBlock)
	for pGame.Fits(pGame.FallingBlock) {
		pGame.FallingBlock.Location.Row++
	}
	pGame.FallingBlock.Location.Row--
	pGame.Put(pGame.FallingBlock)
	pGame.MakeNewBlocks()
}

// Checks whether a block can be placed on the board
func (pGame *Game) Fits(block *Block) bool {
	for i := 0; i < NUM_CELLS; i++ {
		location := Tetrominos[block.BlockType][block.Orientation][i]
		row := block.Location.Row + location.Row
		col := block.Location.Col + location.Col

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
func (pGame *Game) GameOver() bool {

	// Save and restore falling block at the end of this method.
	pGame.Remove(pGame.FallingBlock)
	over := false

	// If the top two rows are not clear, the game is over
	for i := 0; i < 2; i++ {
		for j := 0; j < pGame.NCols; j++ {
			cell := pGame.Get(i, j)
			if cell != TC_EMPTY {
				over = true
			}
		}
	}
	pGame.Put(pGame.FallingBlock)
	return over
}

// Get returns the cell at the given row and column.
func (pGame *Game) Get(row, col int) Cell {
	if ok, _ := pGame.WithinBounds(row, col); !ok {
		return TC_EMPTY
	}
	return pGame.Board[row][col]
}

// HandleMove performs the action specified by the move.
func (pGame *Game) HandleMove(move Move) {
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
func (pGame *Game) Hold() {
	pGame.Remove(pGame.FallingBlock)
	if pGame.StoredBlock == nil {

		// There is no stored block. Just copy the falling one.

		pGame.StoredBlock = pGame.FallingBlock
	} else {

		// There is a stored block. Swap with the falling one.

		typ := pGame.FallingBlock.BlockType
		ori := pGame.FallingBlock.Orientation

		pGame.FallingBlock.BlockType = pGame.StoredBlock.BlockType
		pGame.FallingBlock.Orientation = pGame.StoredBlock.Orientation

		pGame.StoredBlock.BlockType = typ
		pGame.StoredBlock.Orientation = ori
		for !pGame.Fits(pGame.FallingBlock) {
			pGame.FallingBlock.Location.Row--
		}
	}
	pGame.Put(pGame.FallingBlock)
}

// Init initializes the data in a Game object
func (pGame *Game) Init(nRows, nCols int) {

	// Initialize the boilerplate fields
	pGame.NRows = nRows
	pGame.NCols = nCols
	pGame.Board = NewBoard(nRows, nCols)
	pGame.Points = 0
	pGame.Level = 0
	pGame.NextBlock = nil
	pGame.StoredBlock = nil
	pGame.TicksRemaining = GRAVITY_LEVEL[0]
	pGame.LinesRemaining = LINES_PER_LEVEL

	// Initialize the random number generator so that we can generate
	// the first two random tetrominos
	rand.Seed(time.Now().UnixNano())

	// Now run MakeNewBlocks twice to initialize the falling and stored
	// pointers
	pGame.MakeNewBlocks()
	pGame.MakeNewBlocks()
}

// Returns true if line i is full
func (pGame *Game) LineFull(i int) bool {
	for j := 0; j < pGame.NCols; j++ {
		cell := pGame.Get(i, j)
		if cell == TC_EMPTY {
			return false
		}
	}
	return true
}

// MakeNewBlocks moves the next block to the falling block and creates a
// new randomly chosen block
func (pGame *Game) MakeNewBlocks() {
	pGame.FallingBlock = pGame.NextBlock
	pGame.NextBlock = RandomBlock(pGame.NCols)
}

// Move moves the falling tetris block left (-1) or right (+1), given
// that it fits
func (pGame *Game) Move(direction int) {
	pGame.Remove(pGame.FallingBlock)
	pGame.FallingBlock.Location.Col += direction
	if !pGame.Fits(pGame.FallingBlock) {
		pGame.FallingBlock.Location.Col -= direction
	}
	pGame.Put(pGame.FallingBlock)
}

// NewBoard allocates memory for a rows*cols board.
func NewBoard(nRows, nCols int) [][]Cell {

	// Define the board as a slice of slices of Cells
	board := make([][]Cell, nRows)

	// Outer loop creates a row at a time
	for row := range board {
		board[row] = make([]Cell, nCols)
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
func (pGame *Game) Put(block *Block) {
	for i := 0; i < NUM_CELLS; i++ {
		location := Tetrominos[block.BlockType][block.Orientation][i]
		newRow := block.Location.Row + location.Row
		newCol := block.Location.Col + location.Col
		newCell := TypeToCell(block.BlockType)
		pGame.Set(newRow, newCol, newCell)
	}
}

// Clears a block off the board
func (pGame *Game) Remove(block *Block) {
	for i := 0; i < NUM_CELLS; i++ {
		location := Tetrominos[block.BlockType][block.Orientation][i]
		newRow := block.Location.Row + location.Row
		newCol := block.Location.Col + location.Col
		
		pGame.Set(newRow, newCol, TC_EMPTY)
	}
}

// Rotate rotates the falling block in either direction (+/-1), given
// that it fits
func (pGame *Game) Rotate(direction int) {
	pGame.Remove(pGame.FallingBlock)
	for {
		pGame.FallingBlock.Orientation =
			(pGame.FallingBlock.Orientation + direction) % NUM_ORIENTATIONS

		// If the new orientation fits, we're done
		if pGame.Fits(pGame.FallingBlock) {
			break
		}

		// Otherwise, try moving it to the left to make it fit
		pGame.FallingBlock.Location.Col--
		if pGame.Fits(pGame.FallingBlock) {
			break
		}

		// Finally, try moving it to the right to make it fit
		pGame.FallingBlock.Location.Col += 2
		if pGame.Fits(pGame.FallingBlock) {
			break
		}

		// Put it back in its original location and try the next
		// orientation. Worst case, we come back to the original
		// orientation and it fits, so this loop will terminate
		pGame.FallingBlock.Location.Col--
	}
	pGame.Put(pGame.FallingBlock)
}

// Set sets the cell at the given row and column.
func (pGame *Game) Set(row, col int, value Cell) error {
	if _, err := pGame.WithinBounds(row, col); err != nil {
		return err
	}
	pGame.Board[row][col] = value
	return nil
}

// Shift every row above r down one
func (pGame *Game) ShiftLines(r int) {
	for i := r - 1; i >= 0; i-- {
		for j := 0; j < pGame.NCols; j++ {
			cell := pGame.Get(i, j)
			pGame.Set(i+1, j, cell)
			pGame.Set(i, j, TC_EMPTY)
		}
	}
}

// String returns a string representation of the game object
func (pGame *Game) String() string {

	var sb strings.Builder

	sb.WriteString("Game:\n")
	sb.WriteString(fmt.Sprintf("  nRows: %d\n", pGame.NRows))
	sb.WriteString(fmt.Sprintf("  nCols: %d\n", pGame.NCols))
	sb.WriteString("  board: {\n")

	// Write each row
	for _, cells := range pGame.Board {
		sb.WriteString("    {")
		// Write each column in this row
		for _, cell := range cells {
			sb.WriteString(cell.String())
		}
		sb.WriteString("}\n")
	}
	sb.WriteString("  }\n")
	sb.WriteString(fmt.Sprintf("  points: %d\n", pGame.Points))
	sb.WriteString(fmt.Sprintf("  level: %d\n", pGame.Level))

	blockString := func(p *Block) string {
		if p != nil {
			return p.String()
		}
		return "<nil>"
	}
	sb.WriteString(fmt.Sprintf("  falling: %s\n", blockString(pGame.FallingBlock)))
	sb.WriteString(fmt.Sprintf("  next:    %s\n", blockString(pGame.NextBlock)))
	sb.WriteString(fmt.Sprintf("  stored:  %s\n", blockString(pGame.StoredBlock)))
	sb.WriteString(fmt.Sprintf("  ticksRemaining: %d\n", pGame.TicksRemaining))
	sb.WriteString(fmt.Sprintf("  linesRemaining: %d", pGame.LinesRemaining))

	s := sb.String()
	return s
}

// Do a single game tick: process gravity, user input, and score.
// Return true if the game is still running, false if it is over.
func (pGame *Game) Tick(move Move) bool {

	// Handle gravity
	pGame.DoGravityTick()

	// Handle input
	pGame.HandleMove(move)

	// Check for cleared lines
	linesCleared := pGame.CheckLines()
	pGame.AdjustScore(linesCleared)

	// Return whether the game will continue (NOT whether it's over)
	gameOver := pGame.GameOver()
	notGameOver := !gameOver

	return notGameOver

}

// WithinBounds returns an error if the specified row or column is
// contained in the board
func (pGame *Game) WithinBounds(row, col int) (bool, error) {
	switch {
	case row < 0, row >= pGame.NRows:
		errmsg := fmt.Sprintf("row %d is not >= %d and < %d", row, 0, pGame.NRows)
		return false, errors.New(errmsg)
	case col < 0, col >= pGame.NCols:
		errmsg := fmt.Sprintf("col %d is not >= %d and < %d", col, 0, pGame.NCols)
		return false, errors.New(errmsg)
	default:
		return true, nil
	}
}
