"""Tetromino geometry lookup table."""

from __future__ import annotations

from tetris.domain.models import Location, TetrominoType

TETROMINOES: dict[TetrominoType, tuple[tuple[Location, ...], ...]] = {
    TetrominoType.I: (
        (Location(1, 0), Location(1, 1), Location(1, 2), Location(1, 3)),
        (Location(0, 2), Location(1, 2), Location(2, 2), Location(3, 2)),
        (Location(3, 0), Location(3, 1), Location(3, 2), Location(3, 3)),
        (Location(0, 1), Location(1, 1), Location(2, 1), Location(3, 1)),
    ),
    TetrominoType.J: (
        (Location(0, 0), Location(1, 0), Location(1, 1), Location(1, 2)),
        (Location(0, 1), Location(0, 2), Location(1, 1), Location(2, 1)),
        (Location(1, 0), Location(1, 1), Location(1, 2), Location(2, 2)),
        (Location(0, 1), Location(1, 1), Location(2, 0), Location(2, 1)),
    ),
    TetrominoType.L: (
        (Location(0, 2), Location(1, 0), Location(1, 1), Location(1, 2)),
        (Location(0, 1), Location(1, 1), Location(2, 1), Location(2, 2)),
        (Location(1, 0), Location(1, 1), Location(1, 2), Location(2, 0)),
        (Location(0, 0), Location(0, 1), Location(1, 1), Location(2, 1)),
    ),
    TetrominoType.O: (
        (Location(0, 1), Location(0, 2), Location(1, 1), Location(1, 2)),
        (Location(0, 1), Location(0, 2), Location(1, 1), Location(1, 2)),
        (Location(0, 1), Location(0, 2), Location(1, 1), Location(1, 2)),
        (Location(0, 1), Location(0, 2), Location(1, 1), Location(1, 2)),
    ),
    TetrominoType.S: (
        (Location(0, 1), Location(0, 2), Location(1, 0), Location(1, 1)),
        (Location(0, 1), Location(1, 1), Location(1, 2), Location(2, 2)),
        (Location(1, 1), Location(1, 2), Location(2, 0), Location(2, 1)),
        (Location(0, 0), Location(1, 0), Location(1, 1), Location(2, 1)),
    ),
    TetrominoType.T: (
        (Location(0, 1), Location(1, 0), Location(1, 1), Location(1, 2)),
        (Location(0, 1), Location(1, 1), Location(1, 2), Location(2, 1)),
        (Location(1, 0), Location(1, 1), Location(1, 2), Location(2, 1)),
        (Location(0, 1), Location(1, 0), Location(1, 1), Location(2, 1)),
    ),
    TetrominoType.Z: (
        (Location(0, 0), Location(0, 1), Location(1, 1), Location(1, 2)),
        (Location(0, 2), Location(1, 1), Location(1, 2), Location(2, 1)),
        (Location(1, 0), Location(1, 1), Location(2, 1), Location(2, 2)),
        (Location(0, 1), Location(1, 0), Location(1, 1), Location(2, 0)),
    ),
}
