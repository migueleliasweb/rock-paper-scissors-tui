package model

import (
	"fmt"
	"rock-paper-scissors/bubble"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	gameSelectionRockItem    = bubble.SimpleItem{TitleItem: "Rock ✊", DescItem: "Rock blunts Scissors"}
	gameSelectionPaperItem   = bubble.SimpleItem{TitleItem: "Paper ✋", DescItem: "Paper covers Rock"}
	gameSelectionScissorItem = bubble.SimpleItem{TitleItem: "Scissors ✌️", DescItem: "Scissors cuts Paper"}

	player1ListItems = []list.Item{
		gameSelectionRockItem,
		gameSelectionPaperItem,
		gameSelectionScissorItem,
	}

	player2ListItems = []list.Item{
		gameSelectionRockItem,
		gameSelectionPaperItem,
		gameSelectionScissorItem,
	}
)

// focusedState tracks which list is currently active
type focusedState int

const (
	focusLeft focusedState = iota
	focusRight
)

type ModelWithModelAndRounds interface {
	SetGameMode(item list.Item)
	SetGameRounds(item list.Item)
}

// Game holds the application state for the game
type Game struct {
	leftModel   list.Model
	centerModel spinner.Model
	rightModel  *Scoreboard
	focus       focusedState
	width       int
	height      int

	gameMode   list.Item
	gameRounds list.Item
}

func (m *Game) SetGameMode(item list.Item) {
	m.gameMode = item
}

func (m *Game) SetGameRounds(item list.Item) {
	m.gameRounds = item

	if m.rightModel != nil {
		if i, ok := item.(bubble.SimpleItem); ok {
			switch i.TitleItem {
			case "Best of one":
				m.rightModel.RoundsLeft = 1
			case "Best of two":
				m.rightModel.RoundsLeft = 2
			case "Best of three":
				m.rightModel.RoundsLeft = 3
			}
		}
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Game) Init() tea.Cmd {
	m.leftModel = list.New(
		player1ListItems,
		list.NewDefaultDelegate(),
		0,
		0,
	)

	m.leftModel.SetFilteringEnabled(false)
	m.leftModel.SetShowPagination(false)
	m.leftModel.SetShowStatusBar(false)
	m.leftModel.DisableQuitKeybindings()
	m.leftModel.Title = "Player 1"

	m.centerModel = spinner.New()
	m.centerModel.Spinner = spinner.Dot
	m.centerModel.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m.rightModel = &Scoreboard{}

	m.focus = focusLeft

	return m.centerModel.Tick
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle list.Model selections

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Resize each view to fit a third of the screen each
		thirdWidth := m.width/3 - 4
		m.leftModel.SetWidth(thirdWidth)
		m.leftModel.SetHeight(m.height / 2)
	}

	// Update the focused list only
	m.leftModel, cmd = m.leftModel.Update(msg)
	if m.focus == focusLeft {
		cmds = append(cmds, cmd)
	}

	var spinnerCmd tea.Cmd
	m.centerModel, spinnerCmd = m.centerModel.Update(msg)
	cmds = append(cmds, spinnerCmd)

	return m, tea.Batch(cmds...)
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Game) View() string {
	submitButtonStyle := lipgloss.NewStyle().
		MarginTop(1).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#505164"))

	leftList := focusedStyle.Render(m.leftModel.View())
	leftView := lipgloss.JoinVertical(lipgloss.Center, leftList, submitButtonStyle.Render("Submit"))

	thirdWidth := m.width/3 - 4
	halfHeight := m.height / 2

	title := m.leftModel.Styles.Title.Render("NPC")
	centerBoxHeight := halfHeight - lipgloss.Height(title)

	centerContent := fmt.Sprintf(
		"%s\n\n%s",
		m.centerModel.View(),
		lipgloss.NewStyle().Blink(true).Render("Thinking"),
	)

	centerBox := lipgloss.NewStyle().Width(thirdWidth).Height(centerBoxHeight).Align(lipgloss.Center, lipgloss.Center).Render(centerContent)

	centerView := noFocusStyle.Render(lipgloss.JoinVertical(lipgloss.Center, title, centerBox))
	rightView := noFocusStyle.Render(m.rightModel.View())

	// Sets up horizontal layout ("split view")
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftView,
		centerView,
		rightView,
	)
}

// Build-time interface check
var _ tea.Model = &Game{}
var _ ModelWithModelAndRounds = &Game{}
