package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Help struct{}

func NewHelp() *Help {
	return &Help{}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Help) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Help) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Help) View() string {
	helpTextStyle := lipgloss.NewStyle().
		Padding(0).
		Margin(0).
		Foreground(lipgloss.Color("#4A4A4A"))

	return helpTextStyle.Render("Use q/ctrl+c to exit")
}

// Build-time interface checking
var _ tea.Model = &Help{}
