"""Mappings between tetromino types and board cells."""

from tetris.ports.cell import Cell
from tetris.ports.tetromino_type import TetrominoType


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
