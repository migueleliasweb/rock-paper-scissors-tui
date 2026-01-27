package main

import (
	"log"

	"rock-paper-scissors/model"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(&model.App{
		ActiveModel: &model.Welcome{},
		HelpModel:   &model.Help{},
		GameModelBuilder: func(
			gameMode list.Item,
			gameRounds list.Item,
		) tea.Model {
			return model.Game{
				GameMode:   gameMode,
				GameRounds: gameRounds,
			}
		},
	})

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
