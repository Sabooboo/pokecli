package list

import tea "github.com/charmbracelet/bubbletea"

type List struct {
	Pokemon []string // Will be using bubbles list anyways so not a big deal
}

func New() List {
	return List{[]string{"Pikachu"}}
}

func (l List) Init() tea.Cmd {
	return nil
}

func (l List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return l, nil
}

func (l List) View() string {
	return "List stub"
}
