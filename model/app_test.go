package model

import (
	"rock-paper-scissors/bubble"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestAppInit(t *testing.T) {
	app := &App{
		WelcomeModel: &Welcome{},
		GameModel:    &Game{},
		HelpModel:    NewHelp(),
	}

	cmd := app.Init()
	if cmd == nil {
		t.Error("expected Init to return a batch command")
	}

	if app.activeModel != app.WelcomeModel {
		t.Error("expected activeModel to be WelcomeModel initially")
	}
}

func TestAppUpdateQuit(t *testing.T) {
	app := &App{
		WelcomeModel: &Welcome{},
		GameModel:    &Game{},
		HelpModel:    NewHelp(),
	}
	app.Init()

	// Test 'q'
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")}
	updatedModel, cmd := app.Update(msg)
	app = updatedModel.(*App)

	if !app.quitting {
		t.Error("expected quitting to be true after 'q'")
	}

	// Verify tea.Quit is returned (checking if cmd is not nil is a proxy)
	if cmd == nil {
		t.Error("expected tea.Quit command")
	}
}

func TestAppUpdateWindowSize(t *testing.T) {
	app := &App{
		WelcomeModel: &Welcome{},
		GameModel:    &Game{},
		HelpModel:    NewHelp(),
	}
	app.Init()

	msg := tea.WindowSizeMsg{Width: 100, Height: 50}
	updatedModel, _ := app.Update(msg)
	app = updatedModel.(*App)

	if app.windowWidth != 100 || app.windowHeight != 50 {
		t.Errorf("expected window size 100x50, got %dx%d", app.windowWidth, app.windowHeight)
	}
}

func TestAppUpdateWelcomeToGame(t *testing.T) {
	welcome := &Welcome{}
	welcome.Init()
	game := &Game{}
	game.Init()

	app := &App{
		WelcomeModel: welcome,
		GameModel:    game,
		HelpModel:    NewHelp(),
	}
	app.Init()

	// Simulate selection in Welcome model
	// We use the types expected by App.go type assertions
	welcome.SelectedGameMode = bubble.ItemWithDeactivation{TitleItem: "Single Player"}
	welcome.SelectedGameRounds = bubble.SimpleItem{TitleItem: "Best of one"}

	// Send a no-op message to trigger the transition check in App.Update
	// We use a key that Welcome doesn't react to (e.g., 'a')
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("a")}

	updatedModel, cmd := app.Update(msg)
	app = updatedModel.(*App)

	// 1. Check active model switched
	if app.activeModel != app.GameModel {
		t.Error("expected activeModel to switch to GameModel")
	}

	// 2. Check GameModel was configured
	g := app.GameModel.(*Game)
	if g.gameMode == nil {
		t.Error("expected GameModel.gameMode to be set")
	}
	if g.gameRounds == nil {
		t.Error("expected GameModel.gameRounds to be set")
	}

	// 3. Check WelcomeModel selection was reset
	w := app.WelcomeModel.(*Welcome)
	if w.SelectedGameMode != nil || w.SelectedGameRounds != nil {
		t.Error("expected WelcomeModel selections to be reset")
	}

	// 4. Check commands (Init + Resize)
	if cmd == nil {
		t.Error("expected commands returned during transition")
	}
}

func TestAppUpdateRestart(t *testing.T) {
	welcome := &Welcome{}
	game := &Game{}

	app := &App{
		WelcomeModel: welcome,
		GameModel:    game,
		HelpModel:    NewHelp(),
		activeModel:  game, // Start at Game
	}

	msg := RestartGameMsg{}
	updatedModel, cmd := app.Update(msg)
	app = updatedModel.(*App)

	if app.activeModel != app.WelcomeModel {
		t.Error("expected activeModel to switch to WelcomeModel after restart")
	}

	if cmd != nil {
		t.Error("expected no commands after restart")
	}
}

func TestAppView(t *testing.T) {
	app := &App{
		WelcomeModel: &Welcome{},
		GameModel:    &Game{},
		HelpModel:    NewHelp(),
	}
	app.Init()

	view := app.View()
	if view == "" {
		t.Error("expected View to return content")
	}
}
