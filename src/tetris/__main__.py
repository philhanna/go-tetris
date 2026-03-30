"""Package entrypoint for `python -m tetris`."""

from tetris.adapters.terminal_curses import main

if __name__ == "__main__":
    main()
