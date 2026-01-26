package model

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Welcome Displays the welcome page and instructions.
type Welcome struct {
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (model *Welcome) Init() (c tea.Cmd) {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (model *Welcome) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			// Allows for clean exit
			return model, tea.Quit
		}
	}

	return model, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (model *Welcome) View() string {
	vp := viewport.New(30, 5)
	vp.Style.Background(lipgloss.Color("#282a2b"))

	vp.SetContent(`Welcome to the game of
Rock-Paper-Scissors!`)

	return vp.View()
}

// Build-time interface check
var _ tea.Model = &Welcome{}
