package tetris

// TetrisMove is all possible moves to give as input to the game.
type TetrisMove int

const (
	TM_LEFT = iota
	TM_RIGHT
	TM_CLOCK
	TM_COUNTER
	TM_DROP
	TM_HOLD
	TM_NONE
)
