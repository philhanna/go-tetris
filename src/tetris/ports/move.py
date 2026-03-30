"""Move enum definitions."""

from enum import IntEnum


class Move(IntEnum):
    LEFT = 0
    RIGHT = 1
    CLOCK = 2
    COUNTER = 3
    DROP = 4
    HOLD = 5
    NONE = 6
