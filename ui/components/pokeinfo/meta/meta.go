package meta

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/ui/typdef/lang"
	tea "github.com/charmbracelet/bubbletea"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
)

type ability struct {
	desc     string
	isHidden bool
}

type Data struct {
	Common    common.Common
	res       typdef.PokeResult
	entry     int
	types     []apistructs.Type
	desc      string // apistructs.FlavorTextEntries
	abilities []ability
}

func New(info typdef.PokeResult) Data {
	d := Data{
		res:   info,
		entry: info.Pokemon.ID,
		types: info.Types,
	}
	for _, v := range info.Species.FlavorTextEntries {
		if v.Language.Name == lang.English { // Leave room for localisation later
			d.desc = v.FlavorText
			break
		}
	}

	abilities := make([]ability, 0)
	for _, v := range info.Abilities {
		a := ability{
			isHidden: v.IsHidden,
		}

		for _, v := range v.Info.FlavorTextEntries {
			if v.Language.Name == lang.English {
				a.desc = v.FlavorText
				break
			}
		}
		abilities = append(abilities, a)
	}

	d.abilities = abilities

	return d
}

func (d Data) SetSize(width, height int) {
	d.Common.SetSize(width, height)
}

func (d Data) Init() tea.Cmd {
	return nil
}

func (d Data) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return d, nil
}

func (d Data) View() string {
	// Works as of now
	// s := strings.Builder{}
	// s.WriteString(fmt.Sprint(d.abilities, "\n")) // Abilities is nil
	// s.WriteString(fmt.Sprint(d.types[0].Name, typecolor.Get(typecolor.Name(d.types[0].Name)), "\n"))
	// s.WriteString(fmt.Sprint(d.desc, "\n"))
	// return s.String()
	return "Hello"
}
