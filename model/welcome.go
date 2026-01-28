package model

import (
	"rock-paper-scissors/bubble"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	gameModeSinglePlayer = bubble.ItemWithDeactivation{
		TitleItem: "Single Player",
		DescItem:  "Player vs NPC",
	}

	gameModeLocalMultiplayer = bubble.ItemWithDeactivation{
		TitleItem:   "Local Multiplayer (local)",
		DescItem:    "(soon)",
		Deactivated: true,
	}
	gameModeMultiplayer = bubble.ItemWithDeactivation{
		TitleItem:   "Multiplayer",
		DescItem:    "(soon)",
		Deactivated: true,
	}

	gameModeListItems = []list.Item{
		gameModeSinglePlayer,
		gameModeLocalMultiplayer,
		gameModeMultiplayer,
	}

	gameRoundsOne   = bubble.SimpleItem{TitleItem: "Best of one", DescItem: "Single round"}
	gameRoundsThree = bubble.SimpleItem{TitleItem: "Best of three", DescItem: "Three rounds"}
	gameRoundsFive  = bubble.SimpleItem{TitleItem: "Best of five", DescItem: "Five rounds"}

	gameRoundsListItems = []list.Item{
		gameRoundsOne,
		gameRoundsThree,
		gameRoundsFive,
	}
)

// Welcome displays the welcome page.
type Welcome struct {
	SelectedGameMode   list.Item
	SelectedGameRounds list.Item

	gameModeListModel   list.Model
	gameRoundsListModel list.Model

	focusedIndex int
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Welcome) Init() (c tea.Cmd) {

	m.focusedIndex = 0
	m.gameModeListModel = list.New(
		gameModeListItems,
		bubble.DelegateItemWithDeactivation(),
		0,
		0,
	)

	m.gameModeListModel.Title = "Select Game mode"
	m.gameModeListModel.SetFilteringEnabled(false)
	m.gameModeListModel.SetShowPagination(false)
	m.gameModeListModel.SetShowStatusBar(false)
	m.gameModeListModel.DisableQuitKeybindings()

	m.gameRoundsListModel = list.New(
		gameRoundsListItems,
		list.NewDefaultDelegate(),
		0,
		0,
	)

	m.gameRoundsListModel.Title = "Select Rounds"
	m.gameRoundsListModel.SetFilteringEnabled(false)
	m.gameRoundsListModel.SetShowPagination(false)
	m.gameRoundsListModel.SetShowStatusBar(false)
	m.gameRoundsListModel.DisableQuitKeybindings()

	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Welcome) Update(msg tea.Msg) (model tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			if m.focusedIndex == 2 {
				selectedItem := m.gameModeListModel.SelectedItem()
				if item, ok := selectedItem.(bubble.ItemWithDeactivation); ok && item.Deactivated {
					return m, nil
				}

				m.SelectedGameMode = m.gameModeListModel.SelectedItem()
				m.SelectedGameRounds = m.gameRoundsListModel.SelectedItem()
			}

		case "tab":
			m.focusedIndex++

			if m.focusedIndex > 2 {
				m.focusedIndex = 0
			}

		case "shift+tab":
			m.focusedIndex--

			if m.focusedIndex < 0 {
				m.focusedIndex = 2
			}
		}

	case tea.WindowSizeMsg:
		halfWidth := msg.Width/2 - 4
		m.gameModeListModel.SetWidth(halfWidth)
		m.gameModeListModel.SetHeight(msg.Height / 2)
		m.gameRoundsListModel.SetWidth(halfWidth)
		m.gameRoundsListModel.SetHeight(msg.Height / 2)
	}

	var cmd tea.Cmd

	if m.focusedIndex == 0 {
		m.gameModeListModel, cmd = m.gameModeListModel.Update(msg)
	} else if m.focusedIndex == 1 {
		m.gameRoundsListModel, cmd = m.gameRoundsListModel.Update(msg)
	}

	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Welcome) View() string {
	startGameStyle := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#505164"))

	var modeView, roundsView, startGameButtomView string

	if m.focusedIndex == 0 {
		modeView = focusedStyle.Render(m.gameModeListModel.View())
		roundsView = noFocusStyle.Render(m.gameRoundsListModel.View())
		startGameButtomView = startGameStyle.Render("Start Game")
	} else if m.focusedIndex == 1 {
		modeView = noFocusStyle.Render(m.gameModeListModel.View())
		roundsView = focusedStyle.Render(m.gameRoundsListModel.View())
		startGameButtomView = startGameStyle.Render("Start Game")
	} else {
		modeView = noFocusStyle.Render(m.gameModeListModel.View())
		roundsView = noFocusStyle.Render(m.gameRoundsListModel.View())
		startGameButtomView = focusedStyle.Render("Start Game")
	}

	// Sets up horizontal layout ("split view")
	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			modeView,
			roundsView,
		),
		startGameButtomView,
	)
}

// Build-time interface check
var _ tea.Model = &Welcome{}
