package tetris

// Tetromino is a read-only template for a block
type Tetromino struct {
	shape       TetrisType
	orientation int
	location    TetrisLocation
}
