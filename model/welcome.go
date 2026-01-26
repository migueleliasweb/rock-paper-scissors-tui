package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Welcome displays the welcome page.
type Welcome struct{}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Welcome) Init() (c tea.Cmd) {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Welcome) Update(msg tea.Msg) (model tea.Model, c tea.Cmd) {
	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Welcome) View() string {
	style := lipgloss.NewStyle().
		Background(lipgloss.Color("#282a2b")).
		Padding(1, 2)

	text := `Welcome to the game of
Rock-Paper-Scissors!`

	return style.Render(text)
}

// Build-time interface check
var _ tea.Model = &Welcome{}
