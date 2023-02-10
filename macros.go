package main

// This file contains constans and functions that were originally
// C macros

// 2 columns per cell makes the game much nicer.
const (
	COLS_PER_CELL = 2
)

// Cell colors (used in InitColors())
const (
	COLOR_BLACK int16 = iota
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA
	COLOR_CYAN
	COLOR_WHITE
)
