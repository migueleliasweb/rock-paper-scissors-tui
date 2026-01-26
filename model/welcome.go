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
		item{title: "Quit 👋", desc: "Quit"},
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

	m.gameConfig.Title = "Select number of players"
	m.gameConfig.SetFilteringEnabled(false)
	m.gameConfig.SetShowPagination(false)
	m.gameConfig.SetShowStatusBar(false)

	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Welcome) Update(msg tea.Msg) (model tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle list selections

	case tea.WindowSizeMsg:
		halfWidth := msg.Width/2 - 4
		m.gameConfig.SetWidth(halfWidth)
		m.gameConfig.SetHeight(msg.Height / 2)
	}

	cfgModel, cmd := m.gameConfig.Update(msg)
	m.gameConfig = cfgModel

	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Welcome) View() string {
	welcomeTextStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 4).
		MarginTop(0)

	// Sets up horizontal layout ("split view")
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		welcomeTextStyle.Render("Rock-Paper-Scissors Game"),
		focusedStyle.Render(m.gameConfig.View()),
	)
}

// Build-time interface check
var _ tea.Model = &Welcome{}
