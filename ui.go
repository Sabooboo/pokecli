package main

import (
	"fmt"
	"os"

	"github.com/Sabooboo/pokecli/common"
	"github.com/Sabooboo/pokecli/components/selector"

	"github.com/Sabooboo/pokecli/pages/info"
	"github.com/Sabooboo/pokecli/pages/list"
	"github.com/Sabooboo/pokecli/pages/search"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	infoPage = iota
	listPage
	searchPage
)

type UI struct {
	tabs  selector.Selector
	pages []common.Component
}

func initialModel() UI {
	return UI{
		tabs:  selector.New([]string{"Info", "List", "Search"}, 2),
		pages: make([]common.Component, 3),
	}
}

func (ui UI) Init() tea.Cmd {
	ui.pages[infoPage] = info.New()
	ui.pages[listPage] = list.New()
	ui.pages[searchPage] = search.New()
	return nil
}

func (ui UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	m, cmd := ui.tabs.Update(msg)
	ui.tabs = m.(selector.Selector)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	for i, p := range ui.pages {
		m, cmd := p.Update(msg)
		ui.pages[i] = m.(common.Component)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return ui, tea.Quit
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

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("There was a problem running the program:", err)
		os.Exit(1)
	}
}
