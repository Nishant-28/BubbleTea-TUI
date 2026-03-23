package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/Nishant-28/BubbleTea-TUI/store"
)

func main() {
	s, err := store.New("notes.db")
	if err != nil {
		log.Fatal(err)
	}

	m := initialModel(s)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Goodbye!")
}
