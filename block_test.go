package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlock_String(t *testing.T) {
	var block Block
	block.BlockType = TET_J
	block.Orientation = 35
	assert.NotEmpty(t, block)
}