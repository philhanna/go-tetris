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
	sb.WriteString(fmt.Sprintf("  falling: %s\n", blockString(pGame.falling)))
	sb.WriteString(fmt.Sprintf("  next:    %s\n", blockString(pGame.next)))
	sb.WriteString(fmt.Sprintf("  stored:  %s\n", blockString(pGame.stored)))
	sb.WriteString(fmt.Sprintf("  ticksUntilGravity: %d\n", pGame.ticksUntilGravity))
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
	pGame.next = nil
	pGame.stored = nil
	pGame.ticksUntilGravity = GRAVITY_LEVEL[0]
	pGame.linesRemaining = LINES_PER_LEVEL

	// Initialize the random number generator so that we can generate
	// the first two random tetrominos
	rand.Seed(time.Now().UnixNano())

	// Now run NewFalling twice to initialize the falling and stored
	// pointers

	pGame.NewFalling()
	pGame.NewFalling()
}

// NewFalling moves the next block to the falling block and creates a
// new falling block
func (pGame *TetrisGame) NewFalling() {
	pGame.falling = pGame.next
	pGame.next = RandomBlock(pGame.nCols)
}

// Create is the constructor for a TetrisGame object
func Create(nRows, nCols int) TetrisGame {
	pGame := new(TetrisGame)
	pGame.Init(nRows, nCols)
	return *pGame
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
