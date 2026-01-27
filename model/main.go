package model

import (
	"rock-paper-scissors/bubble"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	globalStyle = lipgloss.NewStyle().Margin(1, 2)

	// The style for the active (focused) list
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D5fff6F4"))

	// The style for the inactive list
	noFocusStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder()).
			BorderForeground(lipgloss.Color("#adadad"))
)

// Main controls the app state.
type Main struct {
	ActiveModel tea.Model
	HelpModel   tea.Model
	GameModel   tea.Model
	quitting    bool
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Main) Init() tea.Cmd {
	return tea.Batch(
		m.ActiveModel.Init(),
		m.HelpModel.Init(),
	)
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Global quit handling
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.ActiveModel, cmd = m.ActiveModel.Update(msg)
	cmds = append(cmds, cmd)

	switch m.ActiveModel.(type) {
	case *Welcome:

	default:
		// something is off
	}

	// Check if whether we just selected a game mode
	if welcome, ok := m.ActiveModel.(*Welcome); ok && welcome.selectedGameMode != nil {
		if selected, ok := welcome.selectedGameMode.(bubble.ItemWithDeactivation); ok {
			// Prevent selecting disabled items
			if selected.Disabled {
				welcome.selectedGameMode = nil
				return m, nil
			}

			// Configure and switch to the Game model
			if game, ok := m.GameModel.(*Game); ok {
				game.gameMode = selected.Title()
			}
			m.ActiveModel = m.GameModel
		}
		// Reset the selection so it doesn't trigger again if we return
		welcome.selectedGameMode = nil
	}

	m.HelpModel, cmd = m.HelpModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Main) View() string {
	return globalStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.ActiveModel.View(),
			m.HelpModel.View(),
		),
	)
}
