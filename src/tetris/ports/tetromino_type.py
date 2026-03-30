"""Tetromino type enum definitions."""

from enum import IntEnum


class TetrominoType(IntEnum):
    """Enumeration of the seven tetromino shapes."""

    I = 0
    J = 1
    L = 2
    O = 3
    S = 4
    T = 5
    Z = 6

    def __str__(self) -> str:
        """Return the one-character display representation for the shape."""
        # Match canonical Tetris piece letters.
        return {
            TetrominoType.I: "I",
            TetrominoType.J: "J",
            TetrominoType.L: "L",
            TetrominoType.O: "O",
            TetrominoType.S: "S",
            TetrominoType.T: "T",
            TetrominoType.Z: "Z",
        }.get(self, "?")
