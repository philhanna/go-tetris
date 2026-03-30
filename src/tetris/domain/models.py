"""Core domain value objects and enums."""

from __future__ import annotations

from dataclasses import dataclass
from enum import IntEnum


class Cell(IntEnum):
    EMPTY = 0
    CELLI = 1
    CELLJ = 2
    CELLL = 3
    CELLO = 4
    CELLS = 5
    CELLT = 6
    CELLZ = 7

    def __str__(self) -> str:
        return {
            Cell.EMPTY: ".",
            Cell.CELLI: "I",
            Cell.CELLJ: "J",
            Cell.CELLL: "L",
            Cell.CELLO: "O",
            Cell.CELLS: "S",
            Cell.CELLT: "T",
            Cell.CELLZ: "Z",
        }.get(self, "?")


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


class Move(IntEnum):
    LEFT = 0
    RIGHT = 1
    CLOCK = 2
    COUNTER = 3
    DROP = 4
    HOLD = 5
    NONE = 6


@dataclass
class Location:
    row: int
    col: int


@dataclass
class Block:
    block_type: TetrominoType
    orientation: int
    location: Location


def type_to_cell(typ: TetrominoType) -> Cell:
    return {
        TetrominoType.I: Cell.CELLI,
        TetrominoType.J: Cell.CELLJ,
        TetrominoType.L: Cell.CELLL,
        TetrominoType.O: Cell.CELLO,
        TetrominoType.S: Cell.CELLS,
        TetrominoType.T: Cell.CELLT,
        TetrominoType.Z: Cell.CELLZ,
    }.get(typ, Cell.EMPTY)
