package tetris

import (
	"math/rand"
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
