package tetris

import (
	"testing"
)

func TestBlock_String(t *testing.T) {
	var block Block
	block.BlockType = TET_J
	block.Orientation = 35
}