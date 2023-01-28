package tetris

import (
	"testing"
)

func setup() *TetrisGame {
	tg := new(TetrisGame)
	tg.rows = 20
	tg.cols = 12
	tg.board = make([]TetrisCell, tg.rows*tg.cols)
	for i := 0; i < cap(tg.board); i++ {
		tg.board[i] = TC_EMPTY
	}
	return tg
}

func TestGameGet(t *testing.T) {
	tg := setup()
	// fmt.Printf("tg=%v\n", tg)
	have := tg.Get(0, 0)
	want := TC_EMPTY
	if have != want {
		t.Errorf("have=%d,want=%d", have, want)
	}
}
