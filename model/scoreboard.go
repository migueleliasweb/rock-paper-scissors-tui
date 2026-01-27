package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Scoreboard struct{}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Scoreboard) Init() tea.Cmd {
	panic("not implemented") // TODO: Implement
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Scoreboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	panic("not implemented") // TODO: Implement
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Scoreboard) View() string {
	panic("not implemented") // TODO: Implement
}
