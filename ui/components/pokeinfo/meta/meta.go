package meta

import (
	"fmt"
	"strings"

	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/ui/typdef/lang"
	"github.com/Sabooboo/pokecli/ui/typdef/typecolor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
)

var (
	TitleStyle = lipgloss.NewStyle()

	SubtitleStyle = TitleStyle.Copy().
			Foreground(lipgloss.Color("#111111"))

	TextStyle = TitleStyle.Copy().Border(lipgloss.NormalBorder())

	TypeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
			Padding(0, 1)
	TypeStyles = []lipgloss.Style{TypeStyle.Copy(), TypeStyle.Copy()}
)

type ability struct {
	desc     string
	isHidden bool
}

type Data struct {
	Common    common.Common
	res       typdef.PokeResult
	name      string
	entry     int
	types     []apistructs.Type
	desc      string // apistructs.FlavorTextEntries
	abilities []ability
}

func New(info typdef.PokeResult) Data {
	d := Data{
		res:   info,
		name:  info.Pokemon.Name,
		entry: info.Pokemon.ID,
		types: info.Types,
	}
	for _, v := range info.Species.FlavorTextEntries {
		if v.Language.Name == lang.English { // Leave room for localisation later
			d.desc = strings.Join(strings.Fields(v.FlavorText), " ")
			break
		}
	}

	for _, v := range TypeStyles {
		v.UnsetBackground()
	}
	for i, v := range info.Types {
		TypeStyles[i].Background(
			lipgloss.Color(typecolor.Get(typecolor.Name(v.Name))),
		)
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

func (d Data) SetSize(width, height int) common.Component {
	d.Common.SetSize(width, height)
	TitleStyle.Width(width)
	SubtitleStyle.Width(width)
	TextStyle.Width(width)
	return d
}

func (d Data) Init() tea.Cmd {
	return nil
}

func (d Data) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return d, nil
}

func (d Data) View() string {
	// Works as of now

	// s.WriteString(fmt.Sprint(d.abilities, "\n")) // Abilities is nil
	// s.WriteString(fmt.Sprint(d.types[0].Name, typecolor.Get(typecolor.Name(d.types[0].Name)), "\n"))
	// s.WriteString(fmt.Sprint(d.desc, "\n"))
	// return s.String()

	name := TitleStyle.Render(d.name)
	entry := SubtitleStyle.Render(fmt.Sprint("#", d.entry))
	top := lipgloss.JoinHorizontal(lipgloss.Left, name, " ", entry)
	desc := TextStyle.Render(d.desc)

	types := make([]string, 0, 2)
	for i, v := range d.types {
		types = append(types, TypeStyles[i].Render(v.Name))
	}

	joined := lipgloss.JoinVertical(lipgloss.Left, top, strings.Join(types, " "), desc)

	return joined
}
