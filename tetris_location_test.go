package tetris

import (
	"testing"
)

func TestLocationString(t *testing.T) {
	input := TetrisLocation{21, 14}
	have := input.String()
	want := "{21, 14}"
	if want != have {
		t.Errorf("want=%s, have=%s", want, have)
	}
}

func TestLocationStringNegative(t *testing.T) {
	input := TetrisLocation{-3, 255}
	have := input.String()
	want := "{-3, 255}"
	if want != have {
		t.Errorf("want=%s, have=%s", want, have)
	}
}
