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
		TitleItem: "Local Multiplayer (local)",
		DescItem:  "(soon)",
		Disabled:  true,
	}
	gameModeMultiplayer = bubble.ItemWithDeactivation{
		TitleItem: "Multiplayer",
		DescItem:  "(soon)",
		Disabled:  true,
	}

	gameModeListItems = []list.Item{
		gameModeSinglePlayer,
		gameModeLocalMultiplayer,
		gameModeMultiplayer,
	}

	gameRoundsOne   = simpleItem{title: "Best of one", desc: "Single round"}
	gameRoundsTwo   = simpleItem{title: "Best of two", desc: "Two rounds"}
	gameRoundsThree = simpleItem{title: "Best of three", desc: "Three rounds"}

	gameRoundsListItems = []list.Item{
		gameRoundsOne,
		gameRoundsTwo,
		gameRoundsThree,
	}
)

// Welcome displays the welcome page.
type Welcome struct {
	gameModeListModel list.Model
	selectedGameMode  list.Item

	gameRoundsListModel list.Model
	selectedGameRounds  list.Item
	focusedIndex        int
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Welcome) Init() (c tea.Cmd) {

	m.gameModeListModel = list.New(
		gameModeListItems,
		list.NewDefaultDelegate(),
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
				m.selectedGameMode = m.gameModeListModel.SelectedItem()
				m.selectedGameRounds = m.gameRoundsListModel.SelectedItem()
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
	welcomeTextStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 4)

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
	return lipgloss.JoinVertical(lipgloss.Center,
		welcomeTextStyle.Render("Rock-Paper-Scissors Game"),
		lipgloss.JoinHorizontal(lipgloss.Top, modeView, roundsView),
		startGameButtomView,
	)
}

// Build-time interface check
var _ tea.Model = &Welcome{}
