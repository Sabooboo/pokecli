package pokedisp

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/components/pokeinfo/img"
	"github.com/Sabooboo/pokecli/ui/components/pokeinfo/meta"
	"github.com/Sabooboo/pokecli/ui/components/selector"
	"github.com/Sabooboo/pokecli/ui/typdef"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Pokeinfo includes N components (width/height):
// Image (square, top left): 0.5*width/0.5*height
// Desc including the name, number, types, and dex entry (square, top right): 0.5*width/0.5*height
// Stats represented as a bar chart (square, bottom left): 0.5*width/0.5*height
// Something else in the bottom right?

type Display struct {
	Common   common.Common
	pokemon  typdef.PokeResult
	metaData meta.Data
	image    img.Image
}

func New(pkmn typdef.PokeResult) Display {
	d := Display{
		pokemon:  pkmn,
		Common:   common.Common{},
		metaData: meta.New(pkmn),
		image:    img.New(pkmn),
	}
	return d
}

func (d Display) SetSize(width, height int) common.Component {
	d.Common.SetSize(width, height)
	d.metaData = d.metaData.SetSize(width/2, height-selector.SelectorHeight).(meta.Data)
	d.image = d.image.SetSize(width/2, height-selector.SelectorHeight).(img.Image)
	return d
}

func (d Display) Init() tea.Cmd {
	return nil
}

func (d Display) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return d, nil
}

func (d Display) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Left, d.metaData.View(), d.image.View())
}
