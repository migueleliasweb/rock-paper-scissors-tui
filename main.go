package main

import (
	"log"

	"rock-paper-scissors/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(&model.App{
		WelcomeModel: &model.Welcome{},
		GameModel:    &model.Game{},
		HelpModel:    &model.Help{},
	})

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
