package main

import (
	"log"

	"rock-paper-scissors/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(&model.Welcome{})

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
