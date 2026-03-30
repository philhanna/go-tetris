"""Game aggregate containing all Tetris rules and state transitions."""

from __future__ import annotations

from dataclasses import dataclass, field
from random import Random

from tetris.domain.constants import (
    GRAVITY_LEVEL,
    LINES_PER_LEVEL,
    MAX_LEVEL,
    NUM_CELLS,
    NUM_ORIENTATIONS,
)
from tetris.ports import Block, Cell, Location, Move, TetrominoType, type_to_cell
from tetris.domain.tetrominoes import TETROMINOES


class BoundsError(ValueError):
    """Raised when attempting to access outside of the board."""


def new_board(rows: int, cols: int) -> list[list[Cell]]:
    return [[Cell.EMPTY for _ in range(cols)] for _ in range(rows)]


@dataclass
class Game:
    n_rows: int
    n_cols: int
    rng: Random = field(default_factory=Random)

    board: list[list[Cell]] = field(init=False)
    points: int = field(default=0, init=False)
    level: int = field(default=0, init=False)
    falling_block: Block | None = field(default=None, init=False)
    next_block: Block | None = field(default=None, init=False)
    ticks_remaining: int = field(default=GRAVITY_LEVEL[0], init=False)
    lines_remaining: int = field(default=LINES_PER_LEVEL, init=False)

    def __post_init__(self) -> None:
        self.board = new_board(self.n_rows, self.n_cols)
        self.make_new_blocks()
        self.make_new_blocks()

    def adjust_score(self, lines_cleared: int) -> None:
        line_multiplier = [0, 40, 100, 300, 1200]
        points = line_multiplier[lines_cleared] * (self.level + 1)
        self.points += points

        if lines_cleared >= self.lines_remaining:
            self.level = min(MAX_LEVEL, self.level + 1)
            lines_cleared = self.lines_remaining
            self.lines_remaining = LINES_PER_LEVEL - lines_cleared
        else:
            self.lines_remaining -= lines_cleared

    def check_lines(self) -> int:
        if self.falling_block is None:
            return 0

        self.remove(self.falling_block)

        n_lines = 0
        i = self.n_rows - 1
        while i >= 0:
            if self.line_full(i):
                self.shift_lines(i)
                i += 1
                n_lines += 1
            i -= 1

        self.put(self.falling_block)
        return n_lines

    def do_gravity_tick(self) -> None:
        if self.falling_block is None:
            return

        self.ticks_remaining -= 1
        if self.ticks_remaining <= 0:
            self.remove(self.falling_block)
            self.falling_block.location.row += 1
            if self.fits(self.falling_block):
                self.ticks_remaining = GRAVITY_LEVEL[self.level]
            else:
                self.falling_block.location.row -= 1
                self.put(self.falling_block)
                self.make_new_blocks()
            if self.falling_block is not None:
                self.put(self.falling_block)

    def down(self) -> None:
        if self.falling_block is None:
            return

        self.remove(self.falling_block)
        while self.fits(self.falling_block):
            self.falling_block.location.row += 1
        self.falling_block.location.row -= 1
        self.put(self.falling_block)
        self.make_new_blocks()

    def fits(self, block: Block) -> bool:
        for i in range(NUM_CELLS):
            location = TETROMINOES[block.block_type][block.orientation][i]
            row = block.location.row + location.row
            col = block.location.col + location.col

            if not self.within_bounds(row, col):
                return False
            if self.get(row, col) != Cell.EMPTY:
                return False
        return True

    def game_over(self) -> bool:
        if self.falling_block is None:
            return False

        self.remove(self.falling_block)
        over = False
        for i in range(2):
            for j in range(self.n_cols):
                if self.get(i, j) != Cell.EMPTY:
                    over = True
        self.put(self.falling_block)
        return over

    def get(self, row: int, col: int) -> Cell:
        if not self.within_bounds(row, col):
            return Cell.EMPTY
        return self.board[row][col]

    def handle_move(self, move: Move) -> None:
        if move == Move.LEFT:
            self.move(-1)
        elif move == Move.RIGHT:
            self.move(1)
        elif move == Move.DROP:
            self.down()
        elif move == Move.CLOCK:
            self.rotate(1)
        elif move == Move.COUNTER:
            self.rotate(-1)

    def line_full(self, row: int) -> bool:
        for col in range(self.n_cols):
            if self.get(row, col) == Cell.EMPTY:
                return False
        return True

    def make_new_blocks(self) -> None:
        self.falling_block = self.next_block
        self.next_block = self.random_block(self.n_cols)

    def move(self, direction: int) -> None:
        if self.falling_block is None:
            return

        self.remove(self.falling_block)
        self.falling_block.location.col += direction
        if not self.fits(self.falling_block):
            self.falling_block.location.col -= direction
        self.put(self.falling_block)

    def put(self, block: Block) -> None:
        for i in range(NUM_CELLS):
            location = TETROMINOES[block.block_type][block.orientation][i]
            new_row = block.location.row + location.row
            new_col = block.location.col + location.col
            self.set(new_row, new_col, type_to_cell(block.block_type))

    def remove(self, block: Block) -> None:
        for i in range(NUM_CELLS):
            location = TETROMINOES[block.block_type][block.orientation][i]
            new_row = block.location.row + location.row
            new_col = block.location.col + location.col
            self.set(new_row, new_col, Cell.EMPTY)

    def rotate(self, direction: int) -> None:
        if self.falling_block is None:
            return

        self.remove(self.falling_block)
        while True:
            self.falling_block.orientation = (
                self.falling_block.orientation + direction + NUM_ORIENTATIONS
            ) % NUM_ORIENTATIONS

            if self.fits(self.falling_block):
                break

            self.falling_block.location.col -= 1
            if self.fits(self.falling_block):
                break

            self.falling_block.location.col += 2
            if self.fits(self.falling_block):
                break

            self.falling_block.location.col -= 1
        self.put(self.falling_block)

    def set(self, row: int, col: int, value: Cell) -> None:
        if not self.within_bounds(row, col):
            raise BoundsError(
                f"row {row} is not >= 0 and < {self.n_rows} or col {col} is not >= 0 and < {self.n_cols}"
            )
        self.board[row][col] = value

    def shift_lines(self, row: int) -> None:
        for i in range(row - 1, -1, -1):
            for j in range(self.n_cols):
                self.set(i + 1, j, self.get(i, j))
                self.set(i, j, Cell.EMPTY)

    def tick(self, move: Move) -> bool:
        self.do_gravity_tick()
        self.handle_move(move)
        lines_cleared = self.check_lines()
        self.adjust_score(lines_cleared)
        return not self.game_over()

    def within_bounds(self, row: int, col: int) -> bool:
        return 0 <= row < self.n_rows and 0 <= col < self.n_cols

    def random_block(self, cols: int) -> Block:
        types = [
            TetrominoType.I,
            TetrominoType.J,
            TetrominoType.L,
            TetrominoType.O,
            TetrominoType.S,
            TetrominoType.T,
            TetrominoType.Z,
        ]
        selected = self.rng.choice(types)
        return Block(block_type=selected, orientation=0, location=Location(0, cols // 2 - 2))
