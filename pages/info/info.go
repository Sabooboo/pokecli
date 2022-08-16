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
	Name  string
	desc  string
	stats Stats
}

func New() Info {
	return Info{}
}

func (i *Info) SetPokemon(name string) {
	i.Name = name
}

func (i Info) Init() tea.Cmd {
	return nil
}

func (i Info) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return i, nil
}

func (i Info) View() string {
	return i.Name
}
