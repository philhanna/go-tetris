package tetris

const (
	TC_EMPTY_STR     = " "
	TC_BLOCK_STR     = "\u2588"
	TETRIS           = 4
	NUM_TETROMINOS   = 7
	NUM_ORIENTATIONS = 4
	MAX_LEVEL        = 19
	LINES_PER_LEVEL  = 10
)

// This array stores all necessary information about the cells that are
// filled by each tetromino.  The first index is the type of the
// tetromino (i.e. shape, e.g. I, J, Z, etc.).  The next index is the
// orientation (0-3).  The final array contains 4 tetris_location
// objects, each mapping to an offset from a point on the upper left
// that is the tetromino "origin".
var TETROMINOS = [][][]TetrisLocation{
	// I block
	{
		{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		{{3, 0}, {3, 1}, {3, 2}, {3, 3}},
		{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
	},
	// J block
	{
		{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
		{{0, 1}, {1, 1}, {2, 0}, {2, 1}},
	},
	// L block
	{
		{{0, 2}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 0}},
		{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
	},
	// O block
	{
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
	},
	// S block
	{
		{{0, 1}, {0, 2}, {1, 0}, {1, 1}},
		{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
		{{1, 1}, {1, 2}, {2, 0}, {2, 1}},
		{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
	},
	// T block
	{
		{{0, 1}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {1, 2}, {2, 1}},
		{{1, 0}, {1, 1}, {1, 2}, {2, 1}},
		{{0, 1}, {1, 0}, {1, 1}, {2, 1}},
	},
	// Z block
	{
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		{{0, 2}, {1, 1}, {1, 2}, {2, 1}},
		{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
		{{0, 1}, {1, 0}, {1, 1}, {2, 0}},
	},
}

// This array tells you how many ticks per gravity by level.  Decreases as level
// increases, to add difficulty.
var GRAVITY_LEVEL = []int{
	50, 48, 46, 44, 42, 40, 38, 36, 34, 32,
	30, 28, 26, 24, 22, 20, 16, 12, 8, 4,
}

func MAX(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MIN(x, y int) int {
	if x < y {
		return x
	}
	return y
}
