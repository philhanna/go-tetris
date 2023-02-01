package tetris

// TetrisMove is an enumeration of all possible moves to give as input
// to the game
type TetrisMove uint8

const (
	TM_LEFT TetrisMove = iota
	TM_RIGHT
	TM_CLOCK
	TM_COUNTER
	TM_DROP
	TM_HOLD
	TM_NONE
)
