# Rock Paper Scissors

A classic Rock Paper Scissors game for your terminal, built with Go and Bubble Tea.

## Overview

This project is a Terminal User Interface (TUI) implementation of the timeless game Rock Paper Scissors. It demonstrates how to build interactive CLI applications using the **Bubble Tea** framework in **Golang**.

## Features

- Game modes:
    - Singleplayer
    - Local multiplayer `(soon)`
    - Remote multiplayer `(soon)`
- Number of rounds:
    - Best of one
    - Best of three
    - Best of five

## Built With

- **[Golang](https://go.dev/)**: The programming language used.
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: An extensible TUI framework.
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)**: For styling the terminal UI.
- **[Bubbles](https://github.com/charmbracelet/bubbles)**: For reuseable TUI components.

## Debugging

Use [build-n-run.sh](./hack/build-n-run.sh) to build and run the game. This is needed as we must pass flags to build command so all symbol information is kept.

Due to the fact this is a TUI application, running it directly from VSCODE is not that straigtforward. It ends up being easier to attach a debugger to a running process.

## How to Play

1. Run the application.
2. Use the **Arrow Keys** or **j/k** to select Rock, Paper, or Scissors.
3. Press **Enter** to confirm your choice.
4. View the result and the updated score.
5. Press **q** or **Ctrl+c** to quit the game.

## Notes

For more information see [NOTES](./NOTES.md).

## License

Distributed under the MIT License. See `LICENSE` for more information.