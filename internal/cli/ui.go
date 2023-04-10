package cli

import (
	"github.com/AlexEkdahl/kango/internal/cli/component"
	"github.com/AlexEkdahl/kango/internal/cli/component/footer"
	"github.com/AlexEkdahl/kango/internal/cli/component/header"
	"github.com/AlexEkdahl/kango/internal/cli/keymap"
	"github.com/AlexEkdahl/kango/internal/cli/message"
	"github.com/AlexEkdahl/kango/internal/cli/pages"
	"github.com/AlexEkdahl/kango/internal/cli/pages/create"
	"github.com/AlexEkdahl/kango/internal/cli/pages/dashboard"
	"github.com/AlexEkdahl/kango/internal/cli/pages/info"
	"github.com/AlexEkdahl/kango/internal/cli/pages/kanbanboard"
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type page int

const (
	dashboardPage page = iota
	kanbanPage
	displayModal
	createModal
)

type UI struct {
	header      component.Model
	pages       []pages.Model
	currentPage page
	footer      component.Model
}

func New(opts Options) UI {
	p := []pages.Model{
		dashboard.New(dashboard.DashboardOptions{
			Client: opts.Client,
		}),
		kanbanboard.New(kanbanboard.KanbanOptions{
			Client: opts.Client,
		}),
		info.New(),
		create.New(),
	}
	currentPage := dashboardPage

	return UI{
		header:      header.New(opts.About.Name, opts.About.Version, opts.About.ShortDescription),
		pages:       p,
		currentPage: currentPage,
		footer:      footer.New(p[currentPage]),
	}
}

func (u UI) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	cmds = append(cmds, u.header.Init())

	for i := range u.pages {
		cmds = append(cmds, u.pages[i].Init())
	}

	return tea.Batch(cmds...)
}

func (u UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		x, y := u.margins()

		u.header = u.header.Resize(msg.Width-x, u.header.Height())
		pageX := msg.Width - x
		pageY := msg.Height - (y + u.header.Height() + u.footer.Height())

		for i := range u.pages {
			u.pages[i] = u.pages[i].Resize(pageX, pageY)
		}
	case message.SelectedBoard:
		u.currentPage = kanbanPage

		u.pages[u.currentPage].Update(msg)
		u = u.refreshFooterKeyMap()
		u = u.refreshHeader(msg)
	case message.ShowItem:
		u.currentPage = displayModal
		u.pages[u.currentPage].Update(msg)
		u = u.refreshFooterKeyMap()
	case message.ShowCreateModal:
		u.currentPage = createModal
		u.pages[u.currentPage].Update(msg)
		u = u.refreshFooterKeyMap()
	case message.CreateTask:
		u.currentPage = kanbanPage
		u = u.refreshFooterKeyMap()
	case message.RefreshKeymap:
		u = u.refreshFooterKeyMap()
	case message.CloseDialog:
		u.currentPage = kanbanPage

		u.pages[u.currentPage].Update(msg)
		u = u.refreshFooterKeyMap()
	case message.GoBack:
		if u.currentPage != dashboardPage {
			u.currentPage = dashboardPage
			// u.pages[dashboardPage].Update(msg)

			u = u.refreshFooterKeyMap()
			u = u.resetHeader()
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Quit):
			u.pages[u.currentPage].Update(msg)

			return u, tea.Quit
		}
	}

	var page tea.Model
	page, cmd = u.pages[u.currentPage].Update(msg)
	u.pages[u.currentPage] = page.(pages.Model)
	cmds = append(cmds, cmd)

	return u, tea.Batch(cmds...)
}

func (u UI) View() string {
	view := lipgloss.JoinVertical(
		lipgloss.Left,
		u.header.View(),
		u.pages[u.currentPage].View(),
		u.footer.View(),
	)

	return theme.AppStyle.Render(view)
}

func (u UI) margins() (int, int) {
	s := theme.AppStyle.Copy()
	return s.GetHorizontalFrameSize(), s.GetVerticalFrameSize()
}

func (u UI) refreshFooterKeyMap() UI {
	footer := u.footer.(footer.Model)
	u.footer = footer.SetKeyMap(u.pages[u.currentPage])
	return u
}

func (u UI) refreshHeader(m message.SelectedBoard) UI {
	header := u.header.(header.Model)
	u.header = header.UpdateName(m)
	return u
}

func (u UI) resetHeader() UI {
	header := u.header.(header.Model)
	u.header = header.Reset()
	return u
}
