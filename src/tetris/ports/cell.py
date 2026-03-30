"""Cell enum definitions."""

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
