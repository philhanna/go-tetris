"""Application service that orchestrates one game session."""

from __future__ import annotations

from dataclasses import dataclass

from tetris.application.ports import InputPort, OutputPort, TimingPort
from tetris.domain.game import Game


@dataclass
class RunGameSession:
    game: Game
    input_port: InputPort
    output_port: OutputPort
    timing_port: TimingPort

    def run(self) -> None:
        running = True
        move = self.input_port.read_move()

        while running:
            running = self.game.tick(move)
            self.output_port.render(self.game)
            self.timing_port.sleep_millis(10)
            move = self.input_port.read_move()

        self.output_port.show_game_over(self.game)
