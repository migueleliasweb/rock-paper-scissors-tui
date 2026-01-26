package model

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// The style for the active (focused) list
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")) // Purple

	// The style for the inactive list
	noFocusStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder()).
			BorderForeground(lipgloss.Color("#adadad")) // Grey

	// Global document style
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	listItems = []list.Item{
		item{title: "Rock ✊", desc: "Rock beats Sissors"},
		item{title: "Paper 🤚", desc: "Paper wraps Rock"},
		item{title: "Scissors ✌️", desc: "Scissors cuts Paper"},
	}
)

// item implements the list.Item interface
type item struct {
	title string
		desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// focusedState tracks which list is currently active
type focusedState int

const (
	focusLeft focusedState = iota
	focusRight
)

// Game holds the application state for the game
type Game struct {
	leftList  list.Model
	rightList tea.Model
	focus     focusedState
	quitting  bool
	width     int
	height    int
}

func (m Game) Init() tea.Cmd {
	m.leftList = listItems
	m.rightList = listItems
	m.focus = focusLeft

	return nil
}

func (m Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Resize each panel to fit half the screen each
		halfWidth := m.width/2 - 4
		m.leftList.SetWidth(halfWidth)
		m.rightList.SetWidth(halfWidth)

		m.leftList.SetHeight(m.height - 4)
		m.rightList.SetHeight(m.height - 4)
	}

	// Update the focused list only
	m.leftList, cmd = m.leftList.Update(msg)
	if m.focus == focusLeft {
		cmds = append(cmds, cmd)
	} else {
		m.rightList, cmd = m.rightList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Game) View() string {
	if m.quitting {
		return ""
	}

	var leftView string

	if m.focus == focusLeft {
		leftView = focusedStyle.Render(m.leftList.View())
	} else {
		leftView = noFocusStyle.Render(m.leftList.View())
	}

	var rightView string
	if m.focus == focusRight {
		rightView = focusedStyle.Render(m.rightList.View())
	} else {
		rightView = noFocusStyle.Render(m.rightList.View())
	}

	// Sets up horizontal layout ("split view")
	return docStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			leftView,
			rightView,
		),
	)
}

// func main() {
// 	// Setup List Models
// 	m := Game{
// 		leftList:  list.New(itemsLeft, list.NewDefaultDelegate(), 0, 0),
// 		rightList: list.New(itemsRight, list.NewDefaultDelegate(), 0, 0),
// 		focus:     focusLeft,
// 	}

// 	m.leftList.Title = "Hardware"
// 	m.rightList.Title = "Software"

// 	// Run the Bubble Tea program
// 	p := tea.NewProgram(m, tea.WithAltScreen())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Println("Error running program:", err)
// 		os.Exit(1)
// 	}
// }
