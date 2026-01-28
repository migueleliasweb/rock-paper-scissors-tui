package model

import (
	"rock-paper-scissors/bubble"
	"strings"
	"testing"
)

func TestScoreboardUpdate(t *testing.T) {
	// Initialize with active selections
	m := &Scoreboard{
		LastPlayer1Selection: bubble.SimpleItem{TitleItem: "Rock"},
		LastPlayer2Selection: bubble.SimpleItem{TitleItem: "Paper"},
	}

	// Send ClearSelectionMsg
	updatedModel, cmd := m.Update(ClearSelectionMsg{})

	if cmd != nil {
		t.Error("expected nil command from Update")
	}

	sm, ok := updatedModel.(*Scoreboard)
	if !ok {
		t.Fatal("expected updated model to be *Scoreboard")
	}

	if sm.LastPlayer1Selection != nil {
		t.Error("expected LastPlayer1Selection to be nil after ClearSelectionMsg")
	}
	if sm.LastPlayer2Selection != nil {
		t.Error("expected LastPlayer2Selection to be nil after ClearSelectionMsg")
	}
}

func TestScoreboardView(t *testing.T) {
	m := &Scoreboard{
		Wins:       10,
		Losses:     5,
		RoundsLeft: 3,
	}

	// Test basic stats view
	output := m.View()
	if !strings.Contains(output, "Wins: 10") {
		t.Errorf("expected view to contain 'Wins: 10', got: %s", output)
	}

	// Test view with selections (simulating a round result)
	m.LastPlayer1Selection = bubble.SimpleItem{TitleItem: "Rock"}
	m.LastPlayer2Selection = bubble.SimpleItem{TitleItem: "Paper"}

	output = m.View()
	if !strings.Contains(output, "You: Rock") {
		t.Error("expected view to show player selection")
	}
	if !strings.Contains(output, "NPC: Paper") {
		t.Error("expected view to show NPC selection")
	}
}

func TestScoreboardViewFinal(t *testing.T) {
	tests := []struct {
		name     string
		wins     int
		losses   int
		expected string
	}{
		{"Win", 5, 2, "You Won!"},
		{"Loss", 2, 5, "You Lost!"},
		{"Draw", 3, 3, "It's a Draw!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Scoreboard{Wins: tt.wins, Losses: tt.losses}
			output := m.ViewFinal()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("expected ViewFinal to contain %q, got: %s", tt.expected, output)
			}
		})
	}
}
