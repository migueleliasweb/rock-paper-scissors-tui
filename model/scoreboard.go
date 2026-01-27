package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Scoreboard struct {
	Wins       int
	Losses     int
	RoundsLeft int
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Scoreboard) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Scoreboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Scoreboard) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		fmt.Sprintf("Wins: %d", m.Wins),
		fmt.Sprintf("Losses: %d", m.Losses),
		fmt.Sprintf("Rounds Left: %d", m.RoundsLeft),
	)
}
