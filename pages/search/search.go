package search

import tea "github.com/charmbracelet/bubbletea"

type Search struct {
	Textbox string // Will be using bubbles Search anyways so not a big deal
}

func New() Search {
	return Search{"Example text here..."}
}

func (s Search) Init() tea.Cmd {
	return nil
}

func (s Search) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s Search) View() string {
	return "Search stub"
}
