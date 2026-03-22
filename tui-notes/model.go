package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	message string
}

func initialModel() model {
	return model{
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
