package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Main controls the app state.
type Main struct {
	ActiveModel tea.Model
	HelpModel   tea.Model
	GameModel   tea.Model
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
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	m.ActiveModel, cmd = m.ActiveModel.Update(msg)
	cmds = append(cmds, cmd)

	m.HelpModel, cmd = m.HelpModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Main) View() string {
	return fmt.Sprintf(
		"%s%s",
		m.ActiveModel.View(),
		m.HelpModel.View(),
	)
}
