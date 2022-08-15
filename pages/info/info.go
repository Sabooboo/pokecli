package info

import tea "github.com/charmbracelet/bubbletea"

type Stats struct {
	hp             int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

type Info struct {
	active bool
	name   string
	desc   string
	stats  Stats
}

func New() Info {
	return Info{active: false}
}

func (i Info) Init() tea.Cmd {
	return nil
}

func (i Info) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return i, nil
}

func (i Info) View() string {
	return "Info stub"
}
