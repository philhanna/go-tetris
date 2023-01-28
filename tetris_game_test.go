package tetris

import (
	"log"
	"testing"
)

func setup() *TetrisGame {
	tg := Create(22, 10)
	return tg
}

func dump(label string, tg *TetrisGame) {
	log.Printf("%s\n", label)
	boardString := tg.Dump()
	log.Println(boardString)
}

func TestGameGetSet(t *testing.T) {
	tg := setup()
	tg.Set(3, 4, TC_CELLL)
	dump("After setting (3, 4) to TC_CELLL", tg)
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

func TestTetrisGame_Put(t *testing.T) {
	type fields struct {
		rows                int
		cols                int
		board               []TetrisCell
		points              int
		level               int
		falling             TetrisBlock
		next                TetrisBlock
		stored              TetrisBlock
		ticks_until_gravity int
		linesRemaining      int
	}
	type args struct {
		block TetrisBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"TET_J", fields{rows: 22, cols: 10},
			args{block: TetrisBlock{typ: TET_J, ori: 0, loc: TetrisLocation{4, 5}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tg := Create(22, 10)
			tg.Put(tt.args.block)
			dump("After Put of TET_J", tg)
		})
	}
}
