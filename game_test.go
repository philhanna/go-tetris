package tetris

import (
	"fmt"
	"testing"
)

func TestTetrisGame_String(t *testing.T) {
	pGame := new(TetrisGame)
	fmt.Println(pGame.String())
}
