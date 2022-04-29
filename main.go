package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CrispinStichart/website-via-ssh/blog"
	"github.com/CrispinStichart/website-via-ssh/blog_list"
	tea "github.com/charmbracelet/bubbletea"
)

const BLOG = "/home/critter/programming/CrispinStichart.github.io/_posts/"

const useHighPerformanceRenderer = false

type model struct {
	postsList   blog_list.PostsModel
	currentPost *blog.Model
	ready       bool
	height      int
	width       int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !m.ready {
			// useless?
			m.ready = true
		}

		// useless because we're passing the msg through?
		// won't be useless if we display other components alongside
		// eachother.
		m.width = msg.Width
		m.height = msg.Height
	case blog.GoBackMsg:
		m.currentPost = nil
		m.postsList.Selected = nil
	}

	m.postsList, cmd = m.postsList.Update(msg)

	// We run this after the postslist update, because when we hit enter to select
	// a post, this update function won't run again until an event is triggered (be it
	// a keypress or a resize event). Probably should refactor to use a cmd that returns
	// a tea.Msg.
	selected := m.postsList.Selected
	if selected != nil {
		if m.currentPost == nil {
			b := blog.New(selected, m.height, m.width)
			m.currentPost = &b

		}
		var b blog.Model
		b, cmd = m.currentPost.Update(msg)
		m.currentPost = &b
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n Initializing..."
	}

	if m.currentPost == nil {

		return m.postsList.View()
	}

	return fmt.Sprint(m.currentPost.View())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var ssh = flag.Bool("ssh", false, "Start SSH Server")
	flag.Parse()
	if *ssh {
		startSSH()
	} else {
		startLocal()
	}
}

func startLocal() {
	postsList := blog_list.NewPostsModel(BLOG)

	p := tea.NewProgram(
		model{postsList: postsList},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if err := p.Start(); err != nil {
		fmt.Println("Could not run program:", err)
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
