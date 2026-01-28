package model

import (
	"rock-paper-scissors/bubble"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	globalStyle = lipgloss.NewStyle().Margin(1, 2)

	gameNameTextStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFF")).
				Background(lipgloss.Color("#7D56F4")).
				Padding(1, 2)

	// The style for active (focused) lists
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D5fff6F4"))

	// The style for inactive lists
	noFocusStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder()).
			BorderForeground(lipgloss.Color("#adadad"))
)

// App controls the app state.
type App struct {
	activeModel tea.Model

	WelcomeModel tea.Model
	GameModel    tea.Model
	HelpModel    tea.Model

	quitting bool

	windowWidth  int
	windowHeight int
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *App) Init() tea.Cmd {
	if m.activeModel == nil {
		// Set initial state for the active model
		m.activeModel = m.WelcomeModel
	}

	return tea.Batch(
		m.WelcomeModel.Init(),
		m.GameModel.Init(),
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
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
	}

	m.activeModel, cmd = m.activeModel.Update(msg)
	cmds = append(cmds, cmd)

	// Check if whether we just selected a game mode
	if welcome, ok := m.activeModel.(*Welcome); ok && welcome.SelectedGameMode != nil && welcome.SelectedGameRounds != nil {
		selectedGameMode, okGameMode := welcome.SelectedGameMode.(bubble.ItemWithDeactivation)
		selectedGameRounds, okGameRounds := welcome.SelectedGameRounds.(bubble.SimpleItem)

		if okGameMode && okGameRounds {
			// Initialize the game model to reset state and start the spinner
			cmds = append(cmds, m.GameModel.Init())

			// Configure Game Model with the selected options after checking interface
			switch gameModel := m.GameModel.(type) {
			case ModelWithModelAndRounds:
				gameModel.SetGameMode(selectedGameMode)
				gameModel.SetGameRounds(selectedGameRounds)
			default:
			}

			// Resize the game model to the current window size
			var resizeCmd tea.Cmd
			m.GameModel, resizeCmd = m.GameModel.Update(tea.WindowSizeMsg{
				Width:  m.windowWidth,
				Height: m.windowHeight,
			})
			cmds = append(cmds, resizeCmd)

			// Reset the selection so it doesn't trigger again if we return
			welcome.SelectedGameMode = nil
			welcome.SelectedGameRounds = nil

			// Update active model
			m.activeModel = m.GameModel
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
			lipgloss.Center,
			gameNameTextStyle.Render("Rock-Paper-Scissors Game"),
			m.activeModel.View(),
			m.HelpModel.View(),
		),
	)
}

var _ tea.Model = &App{}
