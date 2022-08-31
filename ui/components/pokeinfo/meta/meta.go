package meta

import (
	"fmt"
	"strings"

	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/ui/typdef/lang"
	"github.com/Sabooboo/pokecli/ui/typdef/typecolor"
	"github.com/Sabooboo/pokecli/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
)

var (
	TitleStyle = lipgloss.NewStyle()

	SubtitleStyle = TitleStyle.Copy().
			Foreground(lipgloss.Color("#111111"))

	DescStyle = TitleStyle.Copy().Border(lipgloss.NormalBorder())

	SubTextStyle = SubtitleStyle.Copy().Italic(true)

	TypeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
			Padding(0, 1)
	TypeStyles = []lipgloss.Style{TypeStyle.Copy(), TypeStyle.Copy()}
)

type ability struct {
	name     string
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
			name:     v.Info.Name,
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
	DescStyle.Width(width)
	return d
}

func (d Data) Init() tea.Cmd {
	return nil
}

func (d Data) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return d, nil
}

func (d Data) View() string {
	name := TitleStyle.Render(util.Title(d.name))
	entry := SubtitleStyle.Render(fmt.Sprint("#", d.entry))
	top := lipgloss.JoinHorizontal(lipgloss.Left, name, " ", entry)
	desc := DescStyle.Render(d.desc)

	types := make([]string, 0, 2)
	for i, v := range d.types {
		types = append(types, TypeStyles[i].Render(v.Name))
	}

	abilities := make([]string, 0)
	for _, v := range d.abilities {
		hiddenFormat := ""
		if v.isHidden {
			hiddenFormat = SubTextStyle.Render("(hidden)")
		}
		top := lipgloss.JoinHorizontal(
			lipgloss.Left,
			TitleStyle.Render(util.Title(v.name)),
			" ",
			hiddenFormat,
		)
		format := lipgloss.JoinVertical(lipgloss.Left, top, TitleStyle.Render(v.desc))
		abilities = append(abilities, format)
	}

	joined := lipgloss.JoinVertical(
		lipgloss.Left,
		top,
		strings.Join(types, " "),
		desc+"\n",
		strings.Join(abilities, "\n\n"),
	)

	return joined
}
