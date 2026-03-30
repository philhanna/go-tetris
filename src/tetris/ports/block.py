"""Block value object."""

from dataclasses import dataclass

from tetris.ports.location import Location
from tetris.ports.tetromino_type import TetrominoType


@dataclass
class Block:
    """Mutable falling piece state: shape, rotation, and board location."""

    block_type: TetrominoType
    orientation: int
    location: Location
