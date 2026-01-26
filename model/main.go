package model

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Main Controls the app state and orchestrates across multiple models and views.
type Main struct {
	Help     help.Model
	KeyMap   help.KeyMap
	ViewPort viewport.Model
	init     bool
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Main) Init() (c tea.Cmd) {
	// Sets the initial state as `init`, so we can present the welcome screen
	m.init = true
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Main) Update(msg tea.Msg) (model tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			// Allows for a clean exit
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Main) View() string {
	vp := viewport.New(30, 30)
	vp.Style.Background(lipgloss.Color("#282a2b"))

	vp.SetContent(`Welcome to the game of
Rock-Paper-Scissors!`)

	return vp.View()
}

// Build-time interface check
var _ tea.Model = &Main{}
