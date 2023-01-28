package tetris

import (
	"testing"
)

func TestCellString(t *testing.T) {
	testCases := []struct {
		cell TetrisCell
		want string
	}{
		{TC_EMPTY, "-"},
		{TC_CELLI, "I"},
		{TC_CELLJ, "J"},
		{TC_CELLL, "L"},
		{TC_CELLO, "O"},
		{TC_CELLS, "S"},
		{TC_CELLT, "T"},
		{TC_CELLZ, "Z"},
		{45, "?"},
	}
	for _, tc := range testCases {
		have := tc.cell.String()
		want := tc.want
		if have != want {
			t.Errorf(`have="%s",want="%s"`, have, want)
		}
	}
}
