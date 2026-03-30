# tests.test_type
from tetris.ports import TetrominoType


def test_type_str() -> None:
    """Verify string conversions for all tetromino enum values."""
    # Ensure enum-to-display conversion remains stable.
    assert str(TetrominoType.I) == "I"
    assert str(TetrominoType.J) == "J"
    assert str(TetrominoType.L) == "L"
    assert str(TetrominoType.O) == "O"
    assert str(TetrominoType.S) == "S"
    assert str(TetrominoType.T) == "T"
    assert str(TetrominoType.Z) == "Z"
