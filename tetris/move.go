package tetris

// Move is an enumeration of all possible moves to give as input
// to the game
type Move uint8

const (
	TM_LEFT Move = iota
	TM_RIGHT
	TM_CLOCK
	TM_COUNTER
	TM_DROP
	TM_HOLD
	TM_NONE
)
