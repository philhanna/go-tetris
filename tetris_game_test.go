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
	type args struct {
		block TetrisBlock
	}
	tests := []struct {
		name   string
		fields TetrisGame
		args   args
	}{
		{"TET_J", *Create(22, 10), args{block: TetrisBlock{typ: TET_J, ori: 0, loc: TetrisLocation{4, 5}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tg := Create(22, 10)
			tg.Put(tt.args.block)
			dump("After Put of TET_J", tg)
		})
	}
}

func TestTetrisGame_Remove(t *testing.T) {
	type args struct {
		block TetrisBlock
	}
	tests := []struct {
		name   string
		fields TetrisGame
		args   args
	}{
		{"TET_J", *Create(22, 10), args{block: TetrisBlock{typ: TET_J, ori: 0, loc: TetrisLocation{4, 5}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tg := Create(22, 10)
			tg.Remove(tt.args.block)
			dump("After Remove of TET_J", tg)
		})
	}
}

func TestTetrisGame_Fits(t *testing.T) {
	type args struct {
		block TetrisBlock
	}
	tests := []struct {
		name   string
		fields TetrisGame
		args   args
		want   bool
	}{
		{
			name:   "TET_J",
			fields: *Create(22, 10),
			args: args{
				block: TetrisBlock{
					typ: TET_J,
					ori: 0,
					loc: TetrisLocation{4, 5}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tg := &TetrisGame{
				rows:                tt.fields.rows,
				cols:                tt.fields.cols,
				board:               tt.fields.board,
				points:              tt.fields.points,
				level:               tt.fields.level,
				falling:             tt.fields.falling,
				next:                tt.fields.next,
				stored:              tt.fields.stored,
				ticks_until_gravity: tt.fields.ticks_until_gravity,
				linesRemaining:      tt.fields.linesRemaining,
			}
			if got := tg.Fits(tt.args.block); got != tt.want {
				t.Errorf("TetrisGame.Fits() = %v, want %v", got, tt.want)
			}
		})
	}
}
