package main

import (
	"log"

	"rock-paper-scissors/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(&model.Main{
		ActiveModel: &model.Welcome{},
		HelpModel:   &model.Help{},
	})

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
