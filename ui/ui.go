package ui

import (
	"fmt"
	"github.com/Sabooboo/pokecli/util"
	"os"
	"time"

	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/Sabooboo/pokecli/ui/components/selector"

	"github.com/Sabooboo/pokecli/ui/pages/info"
	"github.com/Sabooboo/pokecli/ui/pages/list"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	infoPage = iota
	listPage
)

type TickMsg time.Time

// This function is called on Update repeatedly. Without this, things which load
// asynchronously will not be updated in "real time", but instead on other updates
// (key press, window resize on unix, scroll, &c). This returns the current time.
func tickEvery() tea.Cmd {
	return tea.Every(time.Second/5, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type UI struct {
	tabs  selector.Selector
	pages []common.Component
}

func initialModel() UI {
	names := []string{"Info", "Pokemon"}
	return UI{ // TODO: Add settings
		tabs:  selector.New(names, 1),
		pages: make([]common.Component, len(names)),
	}
}

func (ui UI) Init() tea.Cmd {
	ui.pages[infoPage] = info.New()
	ui.pages[listPage] = list.New()
	return tickEvery()
}

func (ui UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	m, cmd := ui.tabs.Update(msg) // Send input to tab selector first
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	old := ui.tabs.Active
	ui.tabs = m.(selector.Selector)
	curr := ui.tabs.Active

	if curr == old { // Tab is the same as last update
		m, cmd := ui.pages[curr].Update(msg)
		ui.pages[curr] = m.(common.Component)
		if cmd != nil {
			if cmd() == list.UpdateMonMsg {
				infoComponent := ui.pages[infoPage].(info.Info)
				selected := ui.pages[listPage].(list.List).Choice
				infoComponent.SetPokemon(selected)
				ui.pages[infoPage] = infoComponent // Update model
				ui.tabs.Active = infoPage
			} else {
				cmds = append(cmds, cmd)
			}

		}
	}

	switch msg := msg.(type) {
	case TickMsg:
		return ui, tickEvery()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return ui, tea.Quit
		}
	case tea.WindowSizeMsg:
		width, height := msg.Width, msg.Height
		selectorHeight := util.Min(height, selector.Height)

		ui.tabs = ui.tabs.SetSize(width, selectorHeight).(selector.Selector)
		for i, v := range ui.pages {
			ui.pages[i] = v.SetSize(width, height-selectorHeight)
		}
	}
	return ui, tea.Batch(cmds...)
}

func (ui UI) View() string {
	var s string
	s += ui.tabs.View() // Tab selector
	s += "\n"
	s += ui.pages[ui.tabs.Active].View() // Active tab
	return s
}

func Start() {
	model := initialModel()
	p := tea.NewProgram(model, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("There was a problem running the program:", err)
		os.Exit(1)
	}
}
