package info

import (
	"github.com/Sabooboo/pokecli/ui/common"
	pokedisp "github.com/Sabooboo/pokecli/ui/components/pokeinfo"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/util"
	tea "github.com/charmbracelet/bubbletea"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
)

var latestRes chan typdef.PokeResult = make(chan typdef.PokeResult, 1)

// type Stats struct {
// 	hp             int
// 	attack         int
// 	defense        int
// 	specialAttack  int
// 	specialDefense int
// 	speed          int
// }

type Info struct {
	Common     common.Common
	pkmn       apistructs.Pokemon
	components pokedisp.Display
	Error      error
	ready      bool
}

func New() Info {
	return Info{ready: false}
}

func (i Info) SetPokemon(name string) {
	go util.GetPokemon(name, latestRes)
}

func (i Info) SetSize(width, height int) common.Component {
	i.Common.SetSize(width, height)
	return i
}

func (i Info) Init() tea.Cmd {
	return nil
}

func (i Info) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(latestRes) > 0 {
		res := <-latestRes

		if res.Error != nil {
			i.Error = res.Error
			return i, nil
		}
		i.components = pokedisp.New(res)
		i.ready = true
		return i, nil
	}
	return i, nil
}

func (i Info) View() string {
	if i.ready {
		return i.components.View()
	}
	return "false"
}
