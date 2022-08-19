package pokedisp

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/components/pokeinfo/meta"
	"github.com/Sabooboo/pokecli/ui/typdef"
	tea "github.com/charmbracelet/bubbletea"
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
}

func New(pkmn typdef.PokeResult) Display {
	d := Display{
		pokemon:  pkmn,
		Common:   common.Common{},
		metaData: meta.New(pkmn),
	}
	return d
}

func (d Display) SetSize(width, height int) {
	d.Common.SetSize(width, height)
}

func (d Display) Init() tea.Cmd {
	return nil
}

func (d Display) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return d, nil
}

func (d Display) View() string {
	return d.metaData.View()
}
