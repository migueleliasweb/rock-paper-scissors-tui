package model

import (
	"fmt"
	"rock-paper-scissors/bubble"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ClearSelectionMsg struct{}

type Scoreboard struct {
	Wins       int
	Losses     int
	Draws      int
	RoundsLeft int

	LastPlayer1Selection list.Item
	LastPlayer2Selection list.Item
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Scoreboard) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Scoreboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case ClearSelectionMsg:
		m.LastPlayer1Selection = nil
		m.LastPlayer2Selection = nil
	}
	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Scoreboard) View() string {
	s := lipgloss.JoinVertical(
		lipgloss.Left,
		fmt.Sprintf("Wins: %d", m.Wins),
		fmt.Sprintf("Losses: %d", m.Losses),
		fmt.Sprintf("Rounds Left: %d", m.RoundsLeft),
	)

	if m.LastPlayer1Selection != nil && m.LastPlayer2Selection != nil {
		p1 := m.LastPlayer1Selection.(bubble.SimpleItem).TitleItem
		p2 := m.LastPlayer2Selection.(bubble.SimpleItem).TitleItem
		s = lipgloss.JoinVertical(lipgloss.Left, s, fmt.Sprintf("\nYou: %s\nNPC: %s", p1, p2))
	}

	return s
}

func (m *Scoreboard) ViewFinal() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		MarginBottom(1)

	var result string
	if m.Wins > m.Losses {
		result = "You Won! 🎉"
	} else if m.Losses > m.Wins {
		result = "You Lost! 💀"
	} else {
		result = "It's a Draw! 🤝"
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render("GAME OVER"),
		lipgloss.NewStyle().MarginBottom(1).Render(result),
		fmt.Sprintf("Wins: %d", m.Wins),
		fmt.Sprintf("Losses: %d", m.Losses),
		fmt.Sprintf("Draws: %d", m.Draws),
	)
}
