"""Tetromino type enum definitions."""

from enum import IntEnum


class TetrominoType(IntEnum):
    I = 0
    J = 1
    L = 2
    O = 3
    S = 4
    T = 5
    Z = 6

    def __str__(self) -> str:
        return {
            TetrominoType.I: "I",
            TetrominoType.J: "J",
            TetrominoType.L: "L",
            TetrominoType.O: "O",
            TetrominoType.S: "S",
            TetrominoType.T: "T",
            TetrominoType.Z: "Z",
        }.get(self, "?")
