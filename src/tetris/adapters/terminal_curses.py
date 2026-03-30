"""Curses-based terminal adapters implementing the application ports."""

from __future__ import annotations

import curses
import time
from dataclasses import dataclass

from tetris.application.ports import InputPort, OutputPort, TimingPort
from tetris.domain.game import Game
from tetris.ports import Block, Cell, Move, TetrominoType, type_to_cell
from tetris.domain.tetrominoes import TETROMINOES

COLS_PER_CELL = 2


@dataclass
class CursesInputAdapter(InputPort):
    stdscr: curses.window

    def read_move(self) -> Move:
        ch = self.stdscr.getch()
        if ch == curses.KEY_LEFT:
            return Move.LEFT
        if ch == curses.KEY_RIGHT:
            return Move.RIGHT
        if ch == curses.KEY_UP:
            return Move.CLOCK
        if ch == curses.KEY_DOWN:
            return Move.COUNTER
        if ch == ord(" "):
            return Move.DROP
        if ch == ord("q"):
            raise KeyboardInterrupt
        return Move.NONE


@dataclass
class CursesTimingAdapter(TimingPort):
    def sleep_millis(self, millis: int) -> None:
        time.sleep(millis / 1000.0)


@dataclass
class CursesOutputAdapter(OutputPort):
    stdscr: curses.window
    board_win: curses.window
    next_win: curses.window
    score_win: curses.window

    def render(self, game: Game) -> None:
        self.display_board(game)
        self.display_piece(game.next_block)
        self.display_score(game)
        curses.doupdate()

    def show_game_over(self, game: Game) -> None:
        self.stdscr.nodelay(False)
        self.stdscr.addstr(game.n_rows + 3, 0, "Game over!\n")
        self.stdscr.addstr(
            game.n_rows + 4,
            0,
            f"You finished with {game.points} points on level {game.level}.\n",
        )
        self.stdscr.getch()

    def display_board(self, game: Game) -> None:
        self.board_win.erase()
        self.board_win.box()
        for i in range(game.n_rows):
            for j in range(game.n_cols):
                cell = game.get(i, j)
                y = 1 + i
                x = 1 + j * COLS_PER_CELL
                if cell != Cell.EMPTY:
                    self.add_block(self.board_win, y, x, cell)
                else:
                    self.add_empty(self.board_win, y, x)
        self.board_win.noutrefresh()

    def display_piece(self, block: Block | None) -> None:
        self.next_win.erase()
        self.next_win.box()
        if block is None:
            self.next_win.noutrefresh()
            return

        for b in range(4):
            location = TETROMINOES[block.block_type][block.orientation][b]
            y = location.row + 1
            x = location.col * COLS_PER_CELL + 1
            self.add_block(self.next_win, y, x, type_to_cell(block.block_type))
        self.next_win.noutrefresh()

    def display_score(self, game: Game) -> None:
        self.score_win.erase()
        self.score_win.box()
        self.score_win.addstr(1, 1, "Score")
        self.score_win.addstr(2, 1, str(game.points))
        self.score_win.addstr(3, 1, "Level")
        self.score_win.addstr(4, 1, str(game.level))
        self.score_win.addstr(5, 1, "Lines")
        self.score_win.addstr(6, 1, str(game.lines_remaining))
        self.score_win.noutrefresh()

    @staticmethod
    def add_block(win: curses.window, y: int, x: int, cell: Cell) -> None:
        color = int(cell)
        try:
            attrs = curses.color_pair(color) | curses.A_REVERSE
        except curses.error:
            attrs = curses.A_REVERSE
        win.addstr(y, x, " " * COLS_PER_CELL, attrs)

    @staticmethod
    def add_empty(win: curses.window, y: int, x: int) -> None:
        win.addstr(y, x, " " * COLS_PER_CELL)


def init_colors() -> None:
    curses.start_color()
    colors = {
        Cell.CELLI: curses.COLOR_CYAN,
        Cell.CELLJ: curses.COLOR_BLUE,
        Cell.CELLL: curses.COLOR_WHITE,
        Cell.CELLO: curses.COLOR_YELLOW,
        Cell.CELLS: curses.COLOR_GREEN,
        Cell.CELLT: curses.COLOR_MAGENTA,
        Cell.CELLZ: curses.COLOR_RED,
    }
    for cell, color in colors.items():
        curses.init_pair(int(cell), color, curses.COLOR_BLACK)


def run_terminal_app(stdscr: curses.window) -> None:
    curses.curs_set(0)
    stdscr.keypad(True)
    stdscr.nodelay(True)

    init_colors()

    game = Game(22, 10)
    h = game.n_rows + 2
    w = COLS_PER_CELL * (game.n_cols + 1)

    board_win = curses.newwin(h, w, 0, 0)
    next_win = curses.newwin(8, 12, 0, w + 3)
    score_win = curses.newwin(9, 12, 12, w + 3)

    output_port = CursesOutputAdapter(stdscr, board_win, next_win, score_win)
    input_port = CursesInputAdapter(stdscr)
    timing_port = CursesTimingAdapter()

    from tetris.application.use_cases import RunGameSession

    use_case = RunGameSession(game, input_port, output_port, timing_port)

    try:
        use_case.run()
    except KeyboardInterrupt:
        pass


def main() -> None:
    curses.wrapper(run_terminal_app)


if __name__ == "__main__":
    main()
