package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/Nishant-28/BubbleTea-TUI/store"
)

type model struct {
	store   store.Store
	notes   []store.Note
	message string
}

func initialModel(s store.Store) model {
	notes, err := s.GetNotes()
	if err != nil {
		notes = []store.Note{}
	}

	return model{
		store:   s,
		notes:   notes,
		message: "Notes loaded.",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("TUI Notes\n")
	b.WriteString("=========\n\n")
	b.WriteString(m.message)
	b.WriteString(fmt.Sprintf("\nTotal notes: %d\n\n", len(m.notes)))

	if len(m.notes) == 0 {
		b.WriteString("No notes found in notes.db yet.\n")
	} else {
		b.WriteString("Saved notes:\n")
		for _, note := range m.notes {
			b.WriteString(fmt.Sprintf("- [%d] %s\n", note.ID, note.Title))
		}
	}

	b.WriteString("\nPress q or ctrl+c to quit.\n")
	return b.String()
}
