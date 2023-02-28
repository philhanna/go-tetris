package tetris

// Number of cells in a tetromino
const NUM_CELLS = 4

// Number of orientations of a tetromino
const NUM_ORIENTATIONS = 4

// Number of tetrominos
const NUM_TETROMINOS = 7

// Highest level
const MAX_LEVEL = 19

// Number of lines per level
const LINES_PER_LEVEL = 10

// Strength of gravity at each level
var GRAVITY_LEVEL = []int{
	50, 48, 46, 44, 42, 40, 38, 36, 34, 32,
	30, 28, 26, 24, 22, 20, 16, 12, 8, 4,
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
