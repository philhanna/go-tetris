package tetris

import (
	"testing"
)

func TestTypeAll(t *testing.T) {
	testCases := []struct {
		typ  TetrisType
		want string
	}{
		{TET_I, "I"},
		{TET_J, "J"},
		{TET_L, "L"},
		{TET_O, "O"},
		{TET_S, "S"},
		{TET_T, "T"},
		{TET_Z, "Z"},
		{86, ""},
	}
	for _, tc := range testCases {
		have := tc.typ.String()
		want := tc.want
		if have != want {
			t.Errorf(`%d: have="%s", want="%s"`, int(tc.typ), have, want)
		}
	}
}
