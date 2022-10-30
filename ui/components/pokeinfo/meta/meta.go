package meta

import (
	"fmt"
	"github.com/Sabooboo/pokecli/ui/components/pokeinfo/statchart"
	"github.com/Sabooboo/pokecli/ui/typdef/poketype"
	"log"
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

	SubtitleStyle = TitleStyle.Copy()

	DescStyle = TitleStyle.Copy().Border(lipgloss.NormalBorder())

	SubTextStyle = SubtitleStyle.Copy().Italic(true)

	TypeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).
			Padding(0, 1)
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
	matchups  poketype.TypeMatchups
	stats     typdef.Stats[int]
	chart     common.Component
	desc      string // apistructs.FlavorTextEntries
	abilities []ability
}

func New(info typdef.PokeResult) Data {
	d := Data{
		res:      info,
		name:     info.Pokemon.Name,
		entry:    info.Pokemon.ID,
		types:    info.Types,
		stats:    info.Stats,
		chart:    statchart.New(info.Stats),
		matchups: poketype.GetTypeMatchups(info.Types...),
	}
	for _, v := range info.Species.FlavorTextEntries {
		if v.Language.Name == lang.English { // Leave room for localisation later
			d.desc = strings.Join(strings.Fields(v.FlavorText), " ")
			break
		}
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

	log.Println(d.stats)
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
	// Top data
	name := TitleStyle.Render(util.Title(d.name))
	entry := SubtitleStyle.Render(fmt.Sprint("#", d.entry))

	// Show name and number next to each other
	top := lipgloss.JoinHorizontal(lipgloss.Left, name, " ", entry)
	desc := DescStyle.Render(d.desc)

	// Types
	types := make([]string, 0, 2)
	for _, v := range d.types {
		types = append(types, GetTypeStyle(typecolor.Name(v.Name)).Render(v.Name))
	}

	// Abilities
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

	// Type Matchups
	immunities := make([]string, 0)
	for _, v := range d.matchups.Immunities {
		immunities = append(immunities, GetTypeStyle(v).Render(string(v)))
	}

	resistances := make([]string, 0)
	for _, v := range d.matchups.Resistances {
		resistances = append(resistances, GetTypeStyle(v).Render(string(v)))
	}

	for _, v := range d.matchups.MajorResistances {
		resistances = append(resistances, GetTypeStyle(v).Render(strings.ToUpper(string(v))))
	}

	weaknesses := make([]string, 0)
	for _, v := range d.matchups.Weaknesses {
		weaknesses = append(weaknesses, GetTypeStyle(v).Render(string(v)))
	}

	for _, v := range d.matchups.MajorWeaknesses {
		weaknesses = append(weaknesses, GetTypeStyle(v).Render(strings.ToUpper(string(v))))
	}

	formattedMatchups := lipgloss.JoinVertical(
		lipgloss.Left,
		"Immunities: "+strings.Join(immunities, ""),
		"Weaknesses: "+strings.Join(weaknesses, ""),
		"Resistances: "+strings.Join(resistances, ""),
	)

	// Display all
	joined := lipgloss.JoinVertical(
		lipgloss.Left,
		top,
		strings.Join(types, " "),
		desc+"\n",
		TitleStyle.Render("Abilities:"),
		strings.Join(abilities, "\n\n"),
		"",
		TitleStyle.Render("Damage Relations:"),
		formattedMatchups,
		"",
		d.chart.View(),
	)

	return joined
}

func GetTypeStyle(name typecolor.Name) lipgloss.Style {
	style := TypeStyle.Copy()
	color := typecolor.Get(name)
	style.Background(lipgloss.Color(color))
	return style
}
