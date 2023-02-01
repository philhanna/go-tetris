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

// MakeNewBlocks moves the next block to the falling block and creates a
// new randomly chosen block
func (pGame *TetrisGame) MakeNewBlocks() {
	pGame.fallingBlock = pGame.nextBlock
	pGame.nextBlock = RandomBlock(pGame.nCols)
}

// WithinBounds returns an error if the specified row or column is
// contained in the board
func (g TetrisGame) WithinBounds(row, col int) (bool, error) {
	switch {
	case row < 0, row >= g.nRows:
		errmsg := fmt.Sprintf("row %d is not >= %d and < %d", row, 0, g.nRows)
		return false, errors.New(errmsg)
	case col < 0, col >= g.nCols:
		errmsg := fmt.Sprintf("col %d is not >= %d and < %d", col, 0, g.nCols)
		return false, errors.New(errmsg)
	default:
		return true, nil
	}
}

// Get returns the cell at the given row and column.
func (g TetrisGame) Get(row, col int) (TetrisCell, error) {
	if ok, err := g.WithinBounds(row, col); !ok {
		return TC_EMPTY, err
	}
	return g.board[row][col], nil
}

// Set sets the cell at the given row and column.
func (g TetrisGame) Set(row, col int, value TetrisCell) error {
	if _, err := g.WithinBounds(row, col); err != nil {
		return err
	}
	g.board[row][col] = value
	return nil
}
