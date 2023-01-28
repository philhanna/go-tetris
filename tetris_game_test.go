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

func TestGameGetSet(t *testing.T) {
	tg := setup()
	tg.Set(3, 4, TC_CELLL)
	have := tg.Get(3, 4)
	want := TC_CELLL
	if have != want {
		t.Errorf("have=%d,want=%d", have, want)
	}
}

func TestGameCheckGood(t *testing.T) {
	tg := setup()
	if !tg.Check(3, 5) {
		t.Error("{3, 5} should have been good")
	}
}

func TestGameCheckBad(t *testing.T) {
	tg := setup()
	if tg.Check(-3, 1000) {
		t.Error("{-3, 1000} should not have been good")
	}
}
