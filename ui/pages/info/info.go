package info

import (
	"github.com/Sabooboo/pokecli/ui/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mtslzr/pokeapi-go"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
)

var latestRes chan pokeResult = make(chan pokeResult, 1)

type pokeResult struct {
	Pokemon apistructs.Pokemon
	Error   error
}

// type Stats struct {
// 	hp             int
// 	attack         int
// 	defense        int
// 	specialAttack  int
// 	specialDefense int
// 	speed          int
// }

type Info struct {
	Common common.Common
	Name   string
	Error  error
	pkmn   apistructs.Pokemon
	ready  bool
}

func New() Info {
	return Info{ready: false}
}

func (i Info) SetPokemon(name string) {
	i.Name = name
	go getPokemon(name)
	i.ready = false
}

func getPokemon(id string) {
	pkmn, err := pokeapi.Pokemon(id)
	latestRes <- pokeResult{
		Pokemon: pkmn,
		Error:   err,
	}
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
		i.pkmn = res.Pokemon
		i.ready = true
		return i, nil
	}
	return i, nil
}

func (i Info) View() string {
	if i.ready {
		return i.pkmn.Abilities[0].Ability.Name
	}
	return "false"
}
