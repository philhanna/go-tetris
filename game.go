package tetris

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// TetrisGame is a class that holds all the information about the state
// of play
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

// String returns a string representation of the game object
func (pGame *TetrisGame) String() string {

	var sb strings.Builder

	sb.WriteString("Game:\n")
	sb.WriteString(fmt.Sprintf("  rows: %d\n", pGame.rows))
	sb.WriteString(fmt.Sprintf("  cols: %d\n", pGame.cols))
	sb.WriteString("  board: {\n")
	// Write each row
	for _, cells := range pGame.board {
		sb.WriteString("    {")
		// Write each column in this row
		for _, cell := range cells {
			sb.WriteString(cell.String())
		}
		sb.WriteString("},\n")
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

// Init initializes the data in a TetrisGame object
func (pGame *TetrisGame) Init(rows, cols int) {

	// Initialize the boilerplate fields
	pGame.rows = rows
	pGame.cols = cols
	pGame.board = make([][]TetrisCell, 0)
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
	pGame.next = RandomBlock(pGame.cols)
}

// Create is the constructor for a TetrisGame object
func Create(rows, cols int) TetrisGame {
	pGame := new(TetrisGame)
	pGame.Init(rows, cols)
	return *pGame
}

// Get returns the cell at the given row and column.
func (g TetrisGame) Get(row, col int) TetrisCell {
	return g.board[row][col]
}

// Set sets the cell at the given row and column.
func (g TetrisGame) Set(row, col int, value TetrisCell) {
	g.board[row][col] = value
}
