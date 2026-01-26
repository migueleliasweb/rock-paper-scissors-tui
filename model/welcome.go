package model

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	configListItems = []list.Item{
		item{title: "Single player ☝️", desc: "Single Player"},
		item{title: "2 players (local) ✌️", desc: "Local Multiplayer"},
	}
)

// Welcome displays the welcome page.
type Welcome struct {
	gameConfig list.Model
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Welcome) Init() (c tea.Cmd) {
	m.gameConfig = list.New(
		configListItems,
		list.NewDefaultDelegate(),
		0,
		0,
	)

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
	welcomeTextStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 4)
		// MarginBottom(1)

	// Sets up horizontal layout ("split view")
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		welcomeTextStyle.Render("Rock-Paper-Scissors Game"),
		focusedStyle.Render(m.gameConfig.View()),
	)
}

// Build-time interface check
var _ tea.Model = &Welcome{}
