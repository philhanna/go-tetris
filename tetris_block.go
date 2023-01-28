package tetris

// TetrisBlock is a struct that contains information about a tetromino.
// Specifically, what type it is, what orientation it has, and where it
// is.
type TetrisBlock struct {
	typ         TetrisType
	orientation int
	location    TetrisLocation
}
