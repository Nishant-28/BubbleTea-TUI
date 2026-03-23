package main

import (
	"fmt"
	"log"

	"github.com/Nishant-28/tui-notes/store"
	tea "github.com/charmbracelet/bubbletea"
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
