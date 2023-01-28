package tetris

import (
	"testing"
)

func setup() *TetrisGame {
	tg := Create(22, 10)
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
