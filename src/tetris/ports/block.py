"""Block value object."""

from dataclasses import dataclass

from tetris.ports.location import Location
from tetris.ports.tetromino_type import TetrominoType


@dataclass
class Block:
    block_type: TetrominoType
    orientation: int
    location: Location
