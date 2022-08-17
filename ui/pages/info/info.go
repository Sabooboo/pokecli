package info

import (
	"github.com/Sabooboo/pokecli/ui/common"
	tea "github.com/charmbracelet/bubbletea"
)

type Stats struct {
	hp             int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

type Info struct {
	Common common.Common
	Name   string
	desc   string
	stats  Stats
}

func New() Info {
	return Info{}
}

func (i Info) SetSize(width, height int) common.Component {
	i.Common.SetSize(width, height)
	return i
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
