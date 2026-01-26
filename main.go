package main

import (
	"log"

	"rock-paper-scissors/model"
	"rock-paper-scissors/view"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	keyMap := &view.KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}

	vp := viewport.New(30, 30)
	vp.Style.Background(lipgloss.Color("#282a2b"))

	vp.SetContent(`Welcome to the game of
Rock-Paper-Scissors!`)

	p := tea.NewProgram(&model.Main{
		ViewPort: viewport.New(30, 30),
		Help:     help.New(),
		KeyMap:   keyMap,
	})

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
