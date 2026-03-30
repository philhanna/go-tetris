"""Location value object."""

from dataclasses import dataclass


@dataclass
class Location:
    """Board coordinate expressed as row and column offsets."""

    row: int
    col: int
