package list

import (
	"fmt"
	"io"

	"github.com/Sabooboo/pokecli/dex"
	"github.com/Sabooboo/pokecli/ui/common"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#ffffff"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	// quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(dex.Pokemon)
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
	Common  common.Common
	list    list.Model
	spinner spinner.Model
	Choice  string
	loading bool
	err     error
}

func New() List {
	s := spinner.New()
	s.Spinner = spinner.Dot

	model := List{loading: false, spinner: s} // TODO: Async loading
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
	var items []list.Item
	nationalPokedex, err := dex.GetPokedex(dex.National, true)
	if err != nil {
		return List{}
	}
	for _, v := range nationalPokedex.Names {
		items = append(items, v)
	}

	const defaultWidth = 40

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose a Pokemon"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Paginator.Type = paginator.Arabic
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	model.list = l
	return model
}

func (l List) SetSize(width, height int) common.Component {
	l.Common.SetSize(width, height)
	if height-10 < listHeight {
		l.list.SetHeight(listHeight)
	} else {
		l.list.SetHeight(height - 10)
	}

	return l
}

func (l List) Init() tea.Cmd {
	return l.spinner.Tick
}

func (l List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		l.list.SetWidth(msg.Width)
		return l, nil
	case tea.KeyMsg:
		l.Choice = ""
		switch msg.String() {
		case "enter":
			item, ok := l.list.SelectedItem().(dex.Pokemon)
			if ok {
				l.Choice = string(item)
			}
			return l, nil
		}
	case List:
		l = msg
	case error:
		l.err = msg
	}
	var cmd tea.Cmd
	l.spinner, cmd = l.spinner.Update(msg)
	cmds = append(cmds, cmd)
	l.list, cmd = l.list.Update(msg)
	cmds = append(cmds, cmd)

	return l, tea.Batch(cmds...)
}

func (l List) View() string {
	if l.loading {
		s := l.spinner.View()
		return fmt.Sprintf("%s Loading... %s", s, s)
	}
	if l.err != nil {
		return l.err.Error()
	}
	return "\n" + l.list.View()
}
