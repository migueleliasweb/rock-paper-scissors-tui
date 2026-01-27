package model

import (
	"rock-paper-scissors/bubble"

	"github.com/charmbracelet/bubbles/list"
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

// App controls the app state.
type App struct {
	ActiveModel      tea.Model
	HelpModel        tea.Model
	GameModelBuilder func(gameMode list.Item, gameRounds list.Item) tea.Model
	quitting         bool
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *App) Init() tea.Cmd {
	return tea.Batch(
		m.ActiveModel.Init(),
		m.HelpModel.Init(),
	)
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	// Check if whether we just selected a game mode
	if welcome, ok := m.ActiveModel.(*Welcome); ok && welcome.SelectedGameMode != nil && welcome.SelectedGameRounds != nil {
		selectedGameMode, okGameMode := welcome.SelectedGameMode.(bubble.ItemWithDeactivation)
		selectedGameRounds, okGameRounds := welcome.SelectedGameRounds.(bubble.ItemWithDeactivation)

		if okGameMode && okGameRounds {
			// Build Game Model with the selected options
			m.ActiveModel = m.GameModelBuilder(
				selectedGameMode,
				selectedGameRounds,
			)

			// Reset the selection so it doesn't trigger again if we return
			// welcome.SelectedGameMode = nil
		}
	}

	m.HelpModel, cmd = m.HelpModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *App) View() string {
	return globalStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.ActiveModel.View(),
			m.HelpModel.View(),
		),
	)
}
