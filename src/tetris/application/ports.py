"""Port definitions for driving and presenting the game."""

from __future__ import annotations

from typing import Protocol

from tetris.domain.game import Game
from tetris.ports import Move


class InputPort(Protocol):
    """Port for reading player intent from any input device."""

    def read_move(self) -> Move:
        """Return the current user move or Move.NONE."""


class OutputPort(Protocol):
    """Port for presenting game state and summary output."""

    def render(self, game: Game) -> None:
        """Render complete game state."""

    def show_game_over(self, game: Game) -> None:
        """Display final score/level and wait for user acknowledgement."""


class TimingPort(Protocol):
    """Port for pacing frame updates."""

    def sleep_millis(self, millis: int) -> None:
        """Sleep for a given time in milliseconds."""
