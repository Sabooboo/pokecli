package selector

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var activeTabBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      " ",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┬",
	TopRight:    "┬",
	BottomLeft:  "┘",
	BottomRight: "└",
}

var tabBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "",
	Right:       "",
	TopLeft:     "─",
	TopRight:    "─",
	BottomLeft:  "─",
	BottomRight: "─",
}

var highlight = lipgloss.AdaptiveColor{Light: "#870000", Dark: "#AA0000"}
var subtle = lipgloss.AdaptiveColor{Light: "#555555", Dark: "#aaaaaa"}

var tabStyleInactive = lipgloss.NewStyle().
	Width(10).
	Align(lipgloss.Center).
	Border(tabBorder).
	BorderForeground(subtle)

var tabStyleActive = tabStyleInactive.Copy().
	Border(activeTabBorder).
	BorderForeground(highlight)

type Selector struct {
	List   []string
	Active int
}

func New(tabs []string, active int) Selector {
	s := Selector{
		List:   make([]string, 0),
		Active: active,
	}
	for _, v := range tabs {
		tab := v
		s.List = append(s.List, tab)
	}
	return s
}

func (s Selector) Init() tea.Cmd {
	return nil
}

func (s Selector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "[", "left":

			if s.Active > 0 {
				s.Active--
			}
		case "]", "right":
			if s.Active < len(s.List)-1 {
				s.Active++
			}
		}
	}
	return s, nil
}

func (s Selector) View() string {
	rendered := make([]string, 0, len(s.List))
	for i, v := range s.List {
		if i == s.Active {
			rendered = append(rendered, tabStyleActive.Render(v))
		} else {
			rendered = append(rendered, tabStyleInactive.Render(v))
		}
	}
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		rendered...,
	)
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, "")
	return row
}