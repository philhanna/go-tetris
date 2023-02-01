package tetris

import (
	"fmt"
	"testing"
)

func getTestGame() *TetrisGame {
	rows, cols := 2, 4
	pGame := TetrisGame{
		rows: rows,
		cols: cols,
	}

	board := make([][]TetrisCell, rows)
	for i := range board {
		board[i] = make([]TetrisCell, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board[row][col] = TC_CELLZ
		}
	}
	// board[0][0] = TC_CELLO
	// board[1][3] = TC_CELLJ
	pGame.board = board

	return &pGame
}

func TestTetrisGame_String(t *testing.T) {
	pGame := getTestGame()
	fmt.Println(pGame.String())
}

func TestTetrisGameGetSet(t *testing.T) {
}
