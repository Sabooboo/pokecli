package pokedisp

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/components/pokeinfo/img"
	"github.com/Sabooboo/pokecli/ui/components/pokeinfo/meta"
	"github.com/Sabooboo/pokecli/ui/typdef"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Display struct {
	Common   common.Common
	pokemon  typdef.PokeResult
	metaData meta.Data
	image    img.Image
}

func New(pkmn typdef.PokeResult, width, height int) Display {
	d := Display{
		pokemon: pkmn,
		Common: common.Common{
			Width:  width,
			Height: height,
		},
		metaData: meta.New(pkmn),
		image:    img.New(pkmn, width/2, height),
	}
	return d
}

func (d Display) SetSize(width, height int) common.Component {
	d.Common.SetSize(width, height)
	d.metaData = d.metaData.SetSize(width/2, height).(meta.Data)
	d.image = d.image.SetSize(width/2, height).(img.Image)
	return d
}

func (d Display) Init() tea.Cmd {
	return nil
}

func (d Display) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	image, _ := d.image.Update(msg)
	d.image = image.(img.Image)
	return d, nil
}

func (d Display) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Left, d.metaData.View(), d.image.View())
}
