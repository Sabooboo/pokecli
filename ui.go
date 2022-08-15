package main

import (
	"fmt"
	"os"

	"github.com/Sabooboo/pokecli/common"
	"github.com/Sabooboo/pokecli/pages/info"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	infoPage = iota
	listPage
	searchPage
)

type UI struct {
	pages  []common.Component
	active int
}

func initialModel() UI {
	return UI{
		pages:  make([]common.Component, 1),
		active: infoPage,
	}
}

func (ui UI) Init() tea.Cmd {
	ui.pages[infoPage] = info.New()
	return nil
}

func (ui UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
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
	return ui.pages[ui.active].View()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("There was a problem running the program:", err)
		os.Exit(1)
	}
}
