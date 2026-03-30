# tests.test_game
from tetris.domain.game import Game, new_board
from tetris.ports import Cell


def get_test_game(n_rows: int, n_cols: int) -> Game:
    """Create a game with a border of occupied cells for fixture-style testing."""
    # Build a deterministic board shape that mimics walls.
    game = Game(n_rows, n_cols)
    game.board = new_board(n_rows, n_cols)
    for col in range(n_cols):
        game.board[0][col] = Cell.CELLO
        game.board[n_rows - 1][col] = Cell.CELLO
    for row in range(1, n_rows - 1):
        game.board[row][0] = Cell.CELLO
        game.board[row][n_cols - 1] = Cell.CELLO
    return game


def test_game_get_set() -> None:
    """Verify that setting a cell can be read back at the same coordinate."""
    # Exercise the basic board read/write API.
    game = get_test_game(22, 10)
    game.set(3, 5, Cell.CELLJ)
    assert game.get(3, 5) == Cell.CELLJ


def test_game_within_bounds() -> None:
    """Verify bounds checking across representative in/out coordinates."""
    # Cover edge and out-of-range coordinates for both axes.
    n_rows = 22
    n_cols = 10
    tests = [
        (12, 5, True),
        (-3, 5, False),
        (0, 5, True),
        (n_rows - 1, 5, True),
        (n_rows, 5, False),
        (n_rows + 17, 5, False),
        (1, -18, False),
        (1, 0, True),
        (1, n_cols - 1, True),
        (1, n_cols, False),
        (1, n_cols + 17, False),
    ]

    for row, col, want in tests:
        game = Game(n_rows, n_cols)
        assert game.within_bounds(row, col) is want
