package tetris

import (
	"fmt"
	"testing"
)

// Creates a TetrisGame with an initialized board
func getTestGame(nRows, nCols int) *TetrisGame {
	pGame := TetrisGame{
		nRows: nRows,
		nCols: nCols,
		board: NewBoard(nRows, nCols),
	}
	for col := 0; col < nCols; col++ {
		pGame.board[0][col] = TC_CELLO
		pGame.board[nRows-1][col] = TC_CELLO
	}
	for row := 1; row < nRows-1; row++ {
		pGame.board[row][0] = TC_CELLO
		pGame.board[row][nCols-1] = TC_CELLO
	}
	return &pGame
}

func TestTetrisPrintGame_String(t *testing.T) {
	t.Skip("Only for debugging")
	pGame := getTestGame(4, 12)
	fmt.Println(pGame.String())
}

func TestTetrisGameGetSet(t *testing.T) {
	pGame := getTestGame(22, 10)
	input := TC_CELLJ
	pGame.Set(3, 5, input)
	output, _ := pGame.Get(3, 5)
	if input != output {
		t.Errorf("Get didn't return what was set: input=%q,output=%q", input, output)
	}
}

func TestTetrisGame_WithinBounds(t *testing.T) {
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
		game := Create(nRows, nCols)
		have, err := game.WithinBounds(tt.row, tt.col)
		if have != tt.want {
			t.Error(err)
		}
	}
}
