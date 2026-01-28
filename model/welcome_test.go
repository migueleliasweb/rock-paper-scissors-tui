package model

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestWelcomeInit(t *testing.T) {
	m := Welcome{}
	m.Init()

	if m.focusedIndex != 0 {
		t.Errorf("expected initial focusedIndex to be 0, got %d", m.focusedIndex)
	}

	if len(m.gameModeListModel.Items()) == 0 {
		t.Error("expected gameModeListModel to be populated")
	}

	if len(m.gameRoundsListModel.Items()) == 0 {
		t.Error("expected gameRoundsListModel to be populated")
	}
}

func TestWelcomeNavigation(t *testing.T) {
	m := Welcome{}
	m.Init()

	// Test Tab Navigation (Forward)
	// 0 (Modes) -> 1 (Rounds) -> 2 (Start Button) -> 0 (Modes)
	tabMsg := tea.KeyMsg{Type: tea.KeyTab}

	// Press Tab -> Focus Rounds (1)
	updatedModel, _ := m.Update(tabMsg)
	m = *updatedModel.(*Welcome)
	if m.focusedIndex != 1 {
		t.Errorf("expected focus 1 after tab, got %d", m.focusedIndex)
	}

	// Press Tab -> Focus Start Button (2)
	updatedModel, _ = m.Update(tabMsg)
	m = *updatedModel.(*Welcome)
	if m.focusedIndex != 2 {
		t.Errorf("expected focus 2 after second tab, got %d", m.focusedIndex)
	}

	// Press Tab -> Loop back to Mode (0)
	updatedModel, _ = m.Update(tabMsg)
	m = *updatedModel.(*Welcome)
	if m.focusedIndex != 0 {
		t.Errorf("expected focus 0 after third tab, got %d", m.focusedIndex)
	}

	// Test Shift+Tab Navigation (Backward)
	// 0 (Modes) -> 2 (Start Button)
	shiftTabMsg := tea.KeyMsg{Type: tea.KeyShiftTab}
	updatedModel, _ = m.Update(shiftTabMsg)
	m = *updatedModel.(*Welcome)
	if m.focusedIndex != 2 {
		t.Errorf("expected focus 2 after shift+tab, got %d", m.focusedIndex)
	}
}

func TestWelcomeSelection(t *testing.T) {
	m := Welcome{}
	m.Init()

	// Move focus to Start Button (index 2)
	m.focusedIndex = 2

	// Ensure valid selections are made in the lists (Default is index 0 for both)
	// Index 0: Single Player (Valid)
	// Index 0: Best of one (Valid)
	m.gameModeListModel.Select(0)
	m.gameRoundsListModel.Select(0)

	enterMsg := tea.KeyMsg{Type: tea.KeyEnter}
	updatedModel, _ := m.Update(enterMsg)
	m = *updatedModel.(*Welcome)

	if m.SelectedGameMode == nil {
		t.Error("expected SelectedGameMode to be set after pressing enter")
	}
	if m.SelectedGameRounds == nil {
		t.Error("expected SelectedGameRounds to be set after pressing enter")
	}

	// Test Deactivated Item Selection
	// Index 1 is "Local Multiplayer" which is Deactivated in welcome.go
	m.gameModeListModel.Select(1)
	m.SelectedGameMode = nil // Reset selection

	updatedModel, _ = m.Update(enterMsg)
	m = *updatedModel.(*Welcome)

	if m.SelectedGameMode != nil {
		t.Error("expected SelectedGameMode to remain nil when selecting a deactivated item")
	}
}
