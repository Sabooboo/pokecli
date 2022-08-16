package list

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mtslzr/pokeapi-go"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#ffffff"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type List struct {
	list   list.Model
	choice string
}

func New() List {
	var items []list.Item
	all, err := pokeapi.Pokedex("national")
	if err != nil {
		return List{}
	}
	for _, v := range all.PokemonEntries {
		items = append(items, item(v.PokemonSpecies.Name))
	}

	const defaultWidth = 40

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose a Pokemon"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	list := List{list: l}

	return list
}

func (l List) Init() tea.Cmd {
	return nil
}

func (l List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		l.list.SetWidth(msg.Width)
		return l, nil
	case tea.KeyMsg:
		l.choice = ""
		switch msg.String() {
		case "enter":
			item, ok := l.list.SelectedItem().(item)
			if ok {
				l.choice = string(item)
			}
			return l, nil
		}
	}

	var cmd tea.Cmd
	l.list, cmd = l.list.Update(msg)
	return l, cmd
}

func (l List) View() string {
	return "\n" + l.list.View()
}
