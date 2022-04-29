package blog_list

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/CrispinStichart/website-via-ssh/formatting"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Post struct {
	Date  time.Time
	title string
	Path  string
}

// tea.Item Interface
func (p Post) Title() string       { return p.title }
func (p Post) Description() string { return formatting.PrettyDate(p.Date) }
func (p Post) FilterValue() string { return p.title }

// stringer interface
func (p Post) String() string {
	return fmt.Sprintf("%v | %v", p.Date, p.title)
}

type Model struct {
	directory string
	posts     list.Model
	Selected  *Post
}

func New(directory string) Model {
	m := Model{
		directory: directory,
		Selected:  nil,
	}

	// figure out how Init is supposed to work and put the IO code there
	m.posts = list.New(getPostsFromDir(directory), list.NewDefaultDelegate(), 0, 0)

	return m
}

func getPostsFromDir(directory string) []list.Item {
	c, err := os.ReadDir(directory)
	check(err)

	posts := make([]list.Item, 0)
	for _, entry := range c {
		if !entry.IsDir() || filepath.Ext(entry.Name()) != ".markdown" {
			date, err := formatting.ExtractDateFromFilename(entry.Name())
			check(err)

			// title, body, err := formatting.SplitTitleFromPost()

			p := Post{
				title: entry.Name(),
				Date:  date,
				Path:  filepath.Join(directory, entry.Name()),
			}
			posts = append(posts, p)
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].(Post).Date.After(posts[j].(Post).Date)
	})

	return posts
}

func (m Model) Init() tea.Cmd {

	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.posts.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			post, ok := m.posts.SelectedItem().(Post)
			if ok {
				m.Selected = &post
			}
			// return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.posts, cmd = m.posts.Update(msg)
	return m, cmd

}

func (m Model) View() string {
	return docStyle.Render(m.posts.View())
}
