package dashboard

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/AlexEkdahl/kango/internal/cli/component/customlist"
	"github.com/AlexEkdahl/kango/internal/cli/keymap"
	"github.com/AlexEkdahl/kango/internal/cli/message"
	"github.com/AlexEkdahl/kango/internal/cli/pages"
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/AlexEkdahl/kango/internal/client"
	"github.com/AlexEkdahl/kango/internal/datastruct"
)

type Model struct {
	viewport   viewport.Model
	boards     list.Model
	styles     *Styles
	loading    spinner.Model
	promptText string
	options    DashboardOptions
}

type DashboardOptions struct {
	Client client.Client
}

type boardsQueriedMsg struct {
	boards *[]datastruct.Board
}

type newBoardMsg struct {
	board *datastruct.Board
}

func New(opt DashboardOptions) Model {
	styles := DefaultStyles()
	loading := spinner.New()
	loading.Spinner = spinner.Dot
	loading.Style = styles.Spinner

	return Model{
		viewport:   viewport.New(0, 0),
		styles:     styles,
		loading:    loading,
		boards:     customlist.New(40, 40),
		promptText: "Select a kanban board:",
		options:    opt,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.viewport.Init(),
		m.loading.Tick,
		func() tea.Msg {
			return m.queryBoards()
		},
		message.RefreshKeyMapCmd,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case boardsQueriedMsg:
		list := []list.Item{}
		boards := msg.boards
		for _, board := range *boards {
			list = append(list, board)
		}

		m.boards.SetItems(list)
		cmds = append(cmds, message.RefreshKeyMapCmd)
	case newBoardMsg:
		board := msg.board
		list := []list.Item{}
		list = append(list, board)

		m.boards.SetItems(list)
		cmds = append(cmds, message.RefreshKeyMapCmd)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Enter):
			if len(m.boards.VisibleItems()) > 1 {
				selected := m.boards.SelectedItem().(datastruct.Board)

				cmds = append(cmds, func() tea.Msg {
					return message.SelectedBoard{
						ID:   selected.ID,
						Name: selected.Name,
						Desc: selected.Desc,
					}
				})
			}
		case key.Matches(msg, keymap.New):
			cmds = append(cmds, m.createBoard)
		case key.Matches(msg, keymap.Quit):
			cmds = append(cmds, tea.Quit)
		}
	}

	m.loading, cmd = m.loading.Update(msg)
	cmds = append(cmds, cmd)

	m.boards, cmd = m.boards.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	list := lipgloss.JoinVertical(
		lipgloss.Top,
		theme.PurpleTextStyle.Render(m.promptText),
		m.boards.View(),
	)

	m.viewport.SetContent(list)

	return m.viewport.View()
}

func (m Model) ShortHelp() []key.Binding {
	kb := make([]key.Binding, 0)

	kb = append(kb, keymap.Quit)

	// Respond to the selection being populated with items
	if len(m.boards.Items()) > 0 {
		if m.boards.FilterState() == list.Filtering {
			kb = append(kb, keymap.Enter, keymap.Escape)
		} else {
			kb = append(kb, keymap.UpDown)

			if m.boards.Paginator.TotalPages > 1 {
				kb = append(kb, keymap.LeftRight)
			}

			kb = append(kb, keymap.Enter, keymap.ForwardSlash)
		}
	}

	return kb
}

func (m Model) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func (m Model) Resize(width, height int) pages.Model {
	m.viewport.Width = width
	m.viewport.Height = height

	return m
}

func (m Model) Width() int {
	return m.viewport.Width
}

func (m Model) Height() int {
	return m.viewport.Height
}

func (m Model) queryBoards() tea.Msg {
	boards, err := m.options.Client.GetAllBoards()
	if err != nil {
		return err
	}

	return boardsQueriedMsg{
		boards: boards,
	}
}

func (m Model) createBoard() tea.Msg {
	board := &datastruct.Board{
		Name: "Kango",
		Desc: "Private",
	}

	id, err := m.options.Client.CreateBoard(board)
	if err != nil {
		return err
	}
	board.ID = *id

	return newBoardMsg{
		board: board,
	}
}
