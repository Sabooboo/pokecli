package search

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Search struct {
	textInput textinput.Model
	err       error
}

// As it turns out, pokeapi has a pokedex call that returns all 898.
// Store in list on init, filterby is search. Perhaps come to this later
// to flesh out in order to search moves, items, locations, etc.
// TODO Improve dynamic searching

func New() Search {
	ti := textinput.New()
	ti.Placeholder = "Staraptor"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 120

	return Search{
		textInput: ti,
		err:       nil,
	}
}

func (s Search) Init() tea.Cmd {
	return textinput.Blink
}

func (s Search) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case error:
		s.err = msg
		return s, nil
	}

	s.textInput, cmd = s.textInput.Update(msg)
	return s, cmd
}

func (s Search) View() string {
	ti := s.textInput
	return fmt.Sprintf(
		"Search pokemon, items, and more\n\n%s\n",
		ti.View(),
	) + "\n"
}
