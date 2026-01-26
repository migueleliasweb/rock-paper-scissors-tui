package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define styles for the layout
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
			BorderForeground(lipgloss.Color("240")) // Grey

	// Global document style
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

// item implements the list.Item interface
type item struct {
	title, desc string
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

// Model holds the application state
type Model struct {
	leftList  list.Model
	rightList list.Model
	focus     focusedState
	quitting  bool
	width     int
	height    int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		// Switch focus with Left/Right arrows
		case "left":
			m.focus = focusLeft
		case "right":
			m.focus = focusRight
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Resize lists to fill half the screen each
		// We subtract some pixels for borders/margins
		halfWidth := m.width/2 - 4
		m.leftList.SetWidth(halfWidth)
		m.rightList.SetWidth(halfWidth)

		m.leftList.SetHeight(m.height - 4)
		m.rightList.SetHeight(m.height - 4)
	}

	// Update the focused list only
	if m.focus == focusLeft {
		m.leftList, cmd = m.leftList.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		m.rightList, cmd = m.rightList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	// 1. Render the Left List
	var leftView string
	if m.focus == focusLeft {
		leftView = focusedStyle.Render(m.leftList.View())
	} else {
		leftView = noFocusStyle.Render(m.leftList.View())
	}

	// 2. Render the Right List
	var rightView string
	if m.focus == focusRight {
		rightView = focusedStyle.Render(m.rightList.View())
	} else {
		rightView = noFocusStyle.Render(m.rightList.View())
	}

	// 3. Join them horizontally
	// lipgloss.JoinHorizontal puts two strings side-by-side
	// lipgloss.Top aligns them to the top
	return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, leftView, rightView))
}

func main() {
	// Initialize Left List items
	itemsLeft := []list.Item{
		item{title: "Raspberry Pi", desc: "Tiny computer"},
		item{title: "Arduino", desc: "Microcontroller"},
		item{title: "ESP32", desc: "WiFi module"},
	}

	// Initialize Right List items
	itemsRight := []list.Item{
		item{title: "Go", desc: "Programming Language"},
		item{title: "Rust", desc: "Programming Language"},
		item{title: "Python", desc: "Programming Language"},
	}

	// Setup List Models
	m := Model{
		leftList:  list.New(itemsLeft, list.NewDefaultDelegate(), 0, 0),
		rightList: list.New(itemsRight, list.NewDefaultDelegate(), 0, 0),
		focus:     focusLeft,
	}

	m.leftList.Title = "Hardware"
	m.rightList.Title = "Software"

	// Run the Bubble Tea program
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
