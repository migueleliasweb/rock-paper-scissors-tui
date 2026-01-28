package model

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestHelpInit(t *testing.T) {
	m := NewHelp()
	cmd := m.Init()
	if cmd != nil {
		t.Error("expected Init to return nil command")
	}
}

func TestHelpUpdate(t *testing.T) {
	m := NewHelp()

	// Send a random message to ensure it's ignored
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("a")}
	updatedModel, cmd := m.Update(msg)

	if cmd != nil {
		t.Error("expected Update to return nil command")
	}

	if updatedModel != m {
		t.Error("expected Update to return the same model instance")
	}
}

func TestHelpView(t *testing.T) {
	m := NewHelp()
	view := m.View()

	expected := "Use q/ctrl+c to exit"
	if !strings.Contains(view, expected) {
		t.Errorf("expected view to contain %q, got %q", expected, view)
	}
}
