package main

import (
	"github.com/Nishant-28/bubble-tea-notes/tui-notes/store"
	tea "github.com/charmbracelet/bubbletea"
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
		message: "hello from Bubble Tea!",
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
	return m.message + "\n\nPress q to quit.\n"
}
