package tetris

import (
	"reflect"
	"testing"
)

func TestRandomTetronimo(t *testing.T) {

	// Test a bunch of random tetrominos and see if they are in the map
	limit := 10
	for i := 0; i < limit; i++ {
		have := RandomTetronimo()
		good := lookup(t, have)
		if !good {
			t.Errorf("Bad tetromino %v\n", have)
		}
	}
	// Now test a bad one, one that isn't in the map
	have := *new(Tetromino)
	good := lookup(t, have)
	if good {
		t.Errorf("Bad tetromino %v\n", have)
	}

}

func lookup(t *testing.T, have Tetromino) bool {
	good := false
	for _, tet := range Tetrominos {
		if reflect.DeepEqual(have, tet) {
			good = true
			break
		}
	}
	return good
}
