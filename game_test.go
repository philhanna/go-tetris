package tetris

import (
	"testing"
)

// Creates a Game with an initialized board
func getTestGame(nRows, nCols int) *Game {
	pGame := Game{
		NRows: nRows,
		NCols: nCols,
		Board: NewBoard(nRows, nCols),
	}
	for col := 0; col < nCols; col++ {
		pGame.Board[0][col] = TC_CELLO
		pGame.Board[nRows-1][col] = TC_CELLO
	}
	for row := 1; row < nRows-1; row++ {
		pGame.Board[row][0] = TC_CELLO
		pGame.Board[row][nCols-1] = TC_CELLO
	}
	return &pGame
}

func TestGameGetSet(t *testing.T) {
	pGame := getTestGame(22, 10)
	input := TC_CELLJ
	pGame.Set(3, 5, input)
	output := pGame.Get(3, 5)
	if input != output {
		t.Errorf("Get didn't return what was set: input=%q,output=%q", input, output)
	}
}

func TestGame_WithinBounds(t *testing.T) {
	nRows := 22
	nCols := 10
	tests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{"good row and column", 12, 5, true},

		{"row way too low", -3, 5, false},
		{"top row", 0, 5, true},
		{"bottom row", nRows - 1, 5, true},
		{"row too high", nRows, 5, false},
		{"row way too high", nRows + 17, 5, false},

		{"col way too far left", 1, -18, false},
		{"left col", 1, 0, true},
		{"right col", 1, nCols - 1, true},
		{"col too far right", 1, nCols, false},
		{"col way too far right", 1, nCols + 17, false},
	}
	for _, tt := range tests {
		game := NewGame(nRows, nCols)
		have, err := game.WithinBounds(tt.row, tt.col)
		if have != tt.want {
			t.Error(err)
		}
	}
}
