package blog

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CrispinStichart/website-via-ssh/blog_list"
	"github.com/CrispinStichart/website-via-ssh/formatting"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	titleStyle = func() lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Italic(true).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		return lipgloss.NewStyle().Align(lipgloss.Right)
	}()
)

type Model struct {
	title    string
	body     string
	date     time.Time
	Height   int
	Width    int
	viewport viewport.Model
}

type KeyMap struct {
	quit key.Binding
	back key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.quit, k.back}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.quit, k.back},
	}
}

var DefaultKeyMap = KeyMap{
	quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
	back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
}

func (m Model) Init() tea.Cmd {
	return nil
}

func readPost(path string) (string, string, error) {
	post, err := os.ReadFile(path)
	check(err)
	return formatting.SplitTitleFromPost(string(post))
}

func New(post *blog_list.Post, height int, width int) Model {
	title, body, err := readPost(post.Path)
	check(err)
	m := Model{
		title: title,
		body:  body,
		date:  post.Date,
	}

	headerHeight := lipgloss.Height(m.headerView())
	footerHeight := lipgloss.Height(m.footerView())
	verticalMarginHeight := headerHeight + footerHeight

	m.viewport = viewport.New(width, height-verticalMarginHeight)
	m.viewport.Width = width
	m.viewport.Height = height - verticalMarginHeight
	m.viewport.YPosition = headerHeight
	// m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
	m.viewport.SetContent(formatting.Glamourize(m.body, width))

	return m
}

type GoBackMsg struct{}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.back):
			return m, func() tea.Msg { return GoBackMsg{} }
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - verticalMarginHeight
		m.viewport.SetContent(formatting.Glamourize(m.body, msg.Width))

	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m Model) headerView() string {
	title := titleStyle.Render(m.title)
	date := " | " + formatting.PrettyDate(m.date)
	return lipgloss.JoinHorizontal(lipgloss.Center, title, date)
}

func (m Model) footerView() string {
	help := help.New().View(DefaultKeyMap)
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	spacer := strings.Repeat(
		" ",
		max(
			0,
			m.viewport.Width-lipgloss.Width(info)-lipgloss.Width(help)))
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		help,
		spacer,
		info,
	)
	// return help.New().View(DefaultKeyMap)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
