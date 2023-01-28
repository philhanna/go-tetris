package main

import (
	"fmt"
	"testing"
)

func TestTetrominos(t *testing.T) {
	for i := range TETROMINOS {
		for j := range TETROMINOS[i] {
			for k := range TETROMINOS[i][j] {
				fmt.Printf("%d %d %d %d\n", i, j, k, TETROMINOS[i][j][k])
			}
		}
	}
}
