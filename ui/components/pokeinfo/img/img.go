package img

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	tea "github.com/charmbracelet/bubbletea"
)

type Image struct {
	common common.Common
}

func New(info typdef.PokeResult) Image {
	i := Image{}
	return i
}

func (i Image) SetSize(width, height int) common.Component {
	i.common.SetSize(width, height)
	return i
}

func (i Image) Init() tea.Cmd {
	return nil
}

func (i Image) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return i, nil
}

func (i Image) View() string {
	return ""
}
