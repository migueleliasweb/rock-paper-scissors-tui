package model

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestGameInit(t *testing.T) {
	m := Game{}
	cmd := m.Init()

	if cmd == nil {
		t.Error("expected Init to return a command (spinner tick)")
	}

	if m.rightModel == nil {
		t.Error("expected rightModel (Scoreboard) to be initialized")
	}
	if m.focus != focusLeft {
		t.Error("expected initial focus to be focusLeft")
	}
	if m.gameOver {
		t.Error("expected gameOver to be false initially")
	}
}

func TestGameUpdateNavigation(t *testing.T) {
	m := Game{}
	m.Init()

	// Initial state: focusLeft
	if m.focus != focusLeft {
		t.Fatalf("expected focusLeft, got %v", m.focus)
	}

	// Press Tab -> focusSubmit
	msg := tea.KeyMsg{Type: tea.KeyTab}
	updatedModel, _ := m.Update(msg)
	m = *updatedModel.(*Game)

	if m.focus != focusSubmit {
		t.Errorf("expected focusSubmit after tab, got %v", m.focus)
	}

	// Press Tab -> focusLeft
	updatedModel, _ = m.Update(msg)
	m = *updatedModel.(*Game)

	if m.focus != focusLeft {
		t.Errorf("expected focusLeft after second tab, got %v", m.focus)
	}
}

func TestGameUpdatePlayRound(t *testing.T) {
	m := Game{}
	m.Init()

	// Setup game with 3 rounds
	m.SetGameRounds(gameRoundsThree)
	initialRounds := m.rightModel.RoundsLeft
	if initialRounds != 3 {
		t.Fatalf("expected 3 rounds, got %d", initialRounds)
	}

	// Select Rock (index 0)
	m.leftModel.Select(0)

	// Move focus to submit button
	m.focus = focusSubmit

	// Press Enter to play
	msg := tea.KeyMsg{Type: tea.KeyEnter}
	updatedModel, cmd := m.Update(msg)
	m = *updatedModel.(*Game)

	// Check if command is returned (ClearSelectionMsg delayed)
	if cmd == nil {
		t.Error("expected command to clear selection")
	}

	// Check rounds decremented
	if m.rightModel.RoundsLeft != 2 {
		t.Errorf("expected 2 rounds left, got %d", m.rightModel.RoundsLeft)
	}

	// Check stats updated (one of them should be > 0)
	stats := m.rightModel.Wins + m.rightModel.Losses + m.rightModel.Draws
	if stats != 1 {
		t.Errorf("expected 1 game played (win/loss/draw), got %d", stats)
	}

	// Check selections are set
	if m.rightModel.LastPlayer1Selection == nil {
		t.Error("expected LastPlayer1Selection to be set")
	}
	if m.rightModel.LastPlayer2Selection == nil {
		t.Error("expected LastPlayer2Selection to be set")
	}
}

func TestGameUpdateGameOver(t *testing.T) {
	m := Game{}
	m.Init()

	// Setup game with 1 round
	m.SetGameRounds(gameRoundsOne)

	// Move focus to submit and play
	m.focus = focusSubmit
	msg := tea.KeyMsg{Type: tea.KeyEnter}

	updatedModel, _ := m.Update(msg)
	m = *updatedModel.(*Game)

	if m.rightModel.RoundsLeft != 0 {
		t.Errorf("expected 0 rounds left, got %d", m.rightModel.RoundsLeft)
	}

	if !m.gameOver {
		t.Error("expected gameOver to be true")
	}

	// Test Restart Trigger
	// Press Enter when game is over
	_, cmd := m.Update(msg)
	if cmd == nil {
		t.Fatal("expected command returned for restart")
	}

	// Execute command to check message type
	msgOut := cmd()
	if _, ok := msgOut.(RestartGameMsg); !ok {
		t.Errorf("expected RestartGameMsg, got %T", msgOut)
	}
}

func TestGameView(t *testing.T) {
	m := Game{}
	m.Init()
	m.width = 100
	m.height = 50

	// Normal view
	view := m.View()
	if !strings.Contains(view, "NPC") {
		t.Error("expected view to contain NPC title")
	}

	// Game Over view
	m.gameOver = true
	view = m.View()
	// ViewFinal contains "GAME OVER" (from scoreboard.go)
	if !strings.Contains(view, "GAME OVER") {
		t.Error("expected view to contain GAME OVER")
	}
}
