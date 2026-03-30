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
    """Reads keyboard input from a curses window and maps to game moves."""

    stdscr: curses.window

    def read_move(self) -> Move:
        """Read one keypress and convert it into a domain Move value."""
        # Map low-level key events to high-level move commands.
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
    """Timing port implementation based on wall-clock sleep."""

    def sleep_millis(self, millis: int) -> None:
        """Pause execution for the requested number of milliseconds."""
        # Convert integer milliseconds into seconds for time.sleep.
        time.sleep(millis / 1000.0)


@dataclass
class CursesOutputAdapter(OutputPort):
    """Renders game state and messages to curses-backed windows."""

    stdscr: curses.window
    board_win: curses.window
    next_win: curses.window
    score_win: curses.window

    def render(self, game: Game) -> None:
        """Render one complete frame of board, preview, and score."""
        # Refresh all UI regions before flushing to the terminal.
        self.display_board(game)
        self.display_piece(game.next_block)
        self.display_score(game)
        curses.doupdate()

    def show_game_over(self, game: Game) -> None:
        """Display final score information and wait for a keypress."""
        # Switch to blocking mode so the summary remains visible.
        self.stdscr.nodelay(False)
        self.stdscr.addstr(game.n_rows + 3, 0, "Game over!\n")
        self.stdscr.addstr(
            game.n_rows + 4,
            0,
            f"You finished with {game.points} points on level {game.level}.\n",
        )
        self.stdscr.getch()

    def display_board(self, game: Game) -> None:
        """Draw the current playfield into the board window."""
        # Iterate row-major and render each board cell as a 2-char block.
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
        """Draw the next-piece preview block."""
        # Clear the preview area first, then draw the tetromino footprint.
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
        """Draw score, level, and remaining lines in the score window."""
        # Keep score labels fixed to avoid visual jitter between frames.
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
        """Render a filled cell at the requested coordinates."""
        # Use reverse-video colored spaces to make blocks look solid.
        color = int(cell)
        try:
            attrs = curses.color_pair(color) | curses.A_REVERSE
        except curses.error:
            attrs = curses.A_REVERSE
        win.addstr(y, x, " " * COLS_PER_CELL, attrs)

    @staticmethod
    def add_empty(win: curses.window, y: int, x: int) -> None:
        """Render an empty cell at the requested coordinates."""
        # Overwrite with plain spaces to clear previous content.
        win.addstr(y, x, " " * COLS_PER_CELL)


def init_colors() -> None:
    """Configure curses color pairs for each tetromino cell type."""
    # Register one color pair per non-empty cell enum value.
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
    """Create terminal adapters and run a full game session."""
    # Initialize terminal behavior for real-time input and drawing.
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
        # Keep control flow in the use case until completion or user abort.
        use_case.run()
    except KeyboardInterrupt:
        pass


def main() -> None:
    """Entrypoint that boots the curses application wrapper."""
    # Let curses manage terminal setup and teardown safely.
    curses.wrapper(run_terminal_app)


if __name__ == "__main__":
    main()
