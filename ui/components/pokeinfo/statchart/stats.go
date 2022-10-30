package statchart

import (
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

const (
	// 255/30 = 8.5, which is a decent height
	maxStat = 255
	ySep    = 30
	block   = "███\n███"
	gap     = "   \n   "
)

type Data struct {
	Common common.Common
	stats  typdef.Stats[int]
	disp   typdef.Stats[[]string]
}

func New(stats typdef.Stats[int]) Data {
	tallest := 0
	createBar := func(stat int) []string {
		bar := make([]string, 0, maxStat/ySep*2)
		for i := 0; i < maxStat/ySep; i++ {
			if i < stat/ySep {
				tallest = util.Max(i, tallest)
				bar = append(bar, block)
			} else {
				bar = append(bar, gap)
			}
		}
		return bar
	}

	eliminateWhitespace := func(bar []string) []string {
		return util.Reverse(bar[0 : tallest+1])
	}

	bars := typdef.Stats[[]string]{
		Health:         createBar(stats.Health),
		Attack:         createBar(stats.Attack),
		SpecialAttack:  createBar(stats.SpecialAttack),
		Defense:        createBar(stats.Defense),
		SpecialDefense: createBar(stats.SpecialDefense),
		Speed:          createBar(stats.Speed),
	}

	data := Data{
		stats: stats,
		disp: typdef.Stats[[]string]{
			Health:         eliminateWhitespace(bars.Health),
			Attack:         eliminateWhitespace(bars.Attack),
			SpecialAttack:  eliminateWhitespace(bars.SpecialAttack),
			Defense:        eliminateWhitespace(bars.Defense),
			SpecialDefense: eliminateWhitespace(bars.SpecialDefense),
			Speed:          eliminateWhitespace(bars.Speed),
		},
	}
	return data
}

func (d Data) SetSize(width, height int) common.Component {
	d.Common.SetSize(width, height)
	return d
}

func (d Data) Init() tea.Cmd {
	return nil
}

func (d Data) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return d, nil
}

func (d Data) View() string {
	return formatBars(d.disp, d.stats)
}

func formatBars(bars typdef.Stats[[]string], vals typdef.Stats[int]) string {
	joinedBars := lipgloss.JoinHorizontal(
		lipgloss.Left,
		formatBar("Hp", vals.Health, bars.Health),
		formatBar("Atk", vals.Attack, bars.Attack),
		formatBar("SpA", vals.SpecialAttack, bars.SpecialAttack),
		formatBar("Def", vals.Defense, bars.Defense),
		formatBar("SpD", vals.SpecialDefense, bars.SpecialDefense),
		formatBar("Spe", vals.Speed, bars.Speed),
	)
	return joinedBars
}

func formatBar(name string, val int, bar []string) string {
	vert := lipgloss.JoinVertical(
		lipgloss.Bottom,
		bar...,
	)
	whole := lipgloss.JoinVertical(
		lipgloss.Left,
		vert,
		name,
		strconv.Itoa(val),
	)
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		whole,
		" ",
	)
}
