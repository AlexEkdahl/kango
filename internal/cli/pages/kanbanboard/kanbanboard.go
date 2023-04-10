package kanbanboard

import (
	"fmt"

	"github.com/AlexEkdahl/kango/internal/cli/component/customlist"
	"github.com/AlexEkdahl/kango/internal/cli/keymap"
	"github.com/AlexEkdahl/kango/internal/cli/message"
	"github.com/AlexEkdahl/kango/internal/cli/pages"
	"github.com/AlexEkdahl/kango/internal/client"
	"github.com/AlexEkdahl/kango/internal/datastruct"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

type KanbanOptions struct {
	Client client.Client
}

type tasksQueriedMsg struct {
	tasks *[]datastruct.Task
}

type Model struct {
	focused  status
	lists    []list.Model
	styles   *Styles
	viewport viewport.Model
	name     string
	options  KanbanOptions
}

func New(opt KanbanOptions) Model {
	lists := make([]list.Model, 3)

	for i := 0; i < 3; i++ {
		lists[i] = customlist.NewBoardList(40, 40)
	}

	lists[todo].Title = "To Do"
	lists[inProgress].Title = "In Progress"
	lists[done].Title = "Done"

	return Model{
		viewport: viewport.New(0, 0),
		styles:   DefaultStyles(),
		focused:  todo,
		lists:    lists,
		options:  opt,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case message.SelectedBoard:
		m.name = msg.Name
		cmds = append(cmds, m.queryTasks)
	case message.CreateTask:

		t, err := m.createTask(&datastruct.Task{
			Status:  datastruct.Status(msg.Status),
			Subject: msg.Title,
			Desc:    msg.Desc,
		})
		if err == nil {
			cmds = append(cmds, m.lists[todo].InsertItem(100, t))
		}
	case tasksQueriedMsg:
		list := []list.Item{}
		for _, task := range *msg.tasks {
			list = append(list, task)
		}

		m.lists[todo].SetItems(list)
		cmds = append(cmds, message.RefreshKeyMapCmd)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Quit):
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, keymap.Enter):
			if len(m.lists[m.focused].VisibleItems()) > 0 {
				selected := m.lists[m.focused].SelectedItem().(datastruct.Task)

				cmds = append(cmds, func() tea.Msg {
					return message.ShowItem{
						ID:    selected.ID,
						Title: selected.Title(),
						Desc:  selected.Description(),
					}
				})
			}
		case key.Matches(msg, keymap.Right):
			if m.focused == done {
				m.focused = todo
			} else {
				m.focused++
			}
		case key.Matches(msg, keymap.Left):
			if m.focused == todo {
				m.focused = done
			} else {
				m.focused--
			}
		case key.Matches(msg, keymap.Escape):
			cmds = append(cmds, func() tea.Msg {
				return message.GoBack{}
			})
		case key.Matches(msg, keymap.Delete):
			if len(m.lists[m.focused].VisibleItems()) > 0 {
				selected := m.lists[m.focused].SelectedItem().(datastruct.Task)
				err := m.delteTask(selected.ID)
				if err == nil {
					m.lists[selected.Status].RemoveItem(m.lists[m.focused].Index())
					cmds = append(cmds, message.RefreshKeyMapCmd)
				}
			}
		case key.Matches(msg, keymap.New):
			cmds = append(cmds, message.ShowCreateModalCmd)
		}
	}

	var list list.Model
	list, cmd = m.lists[m.focused].Update(msg)
	m.lists[m.focused] = list
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var view string
	columnStyle := m.styles.ColumnStyle.Height(m.viewport.Height - 4).Width((m.viewport.Width / 3) - 2)

	todoView := m.lists[todo]
	inProgView := m.lists[inProgress]
	doneView := m.lists[done]

	switch m.focused {
	case inProgress:
		inProgView.SetDelegate(customlist.NewStyledDeligate())

		view = lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(todoView.View()),
			columnStyle.Render(inProgView.View()),
			columnStyle.Render(doneView.View()),
		)
	case done:
		doneView.SetDelegate(customlist.NewStyledDeligate())

		view = lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(todoView.View()),
			columnStyle.Render(inProgView.View()),
			columnStyle.Render(doneView.View()),
		)
	default:
		todoView.SetDelegate(customlist.NewStyledDeligate())

		view = lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(todoView.View()),
			columnStyle.Render(inProgView.View()),
			columnStyle.Render(doneView.View()),
		)
	}

	m.viewport.SetContent(view)

	return m.viewport.View()
}

func (m Model) ShortHelp() []key.Binding {
	kb := make([]key.Binding, 0)

	kb = append(kb, keymap.Quit)

	for _, l := range m.lists {
		if len(l.Items()) > 0 {
			if l.FilterState() == list.Filtering {
				keymap.New.SetEnabled(false)
				kb = append(kb, keymap.Enter, keymap.Escape)
			} else {
				kb = append(kb, keymap.UpDown)
				kb = append(kb, keymap.Delete)
				kb = append(kb, keymap.Enter, keymap.ForwardSlash, keymap.LeftRight, keymap.New)
				keymap.New.SetEnabled(true)
			}
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
	for i, list := range m.lists {
		list.SetSize(width/3, height-4)
		m.lists[i], _ = list.Update("")
	}

	return m
}

func (m Model) Width() int {
	return m.viewport.Width
}

func (m Model) Height() int {
	return m.viewport.Height
}

func (m Model) queryTasks() tea.Msg {
	tasks, err := m.options.Client.GetAllTasks()
	if err != nil {
		return err
	}
	return tasksQueriedMsg{
		tasks: tasks,
	}
}

func (m Model) createTask(task *datastruct.Task) (*datastruct.Task, error) {
	id, err := m.options.Client.CreateTask(task)
	if err != nil {
		return nil, err
	}
	task.ID = *id

	return task, err
}

func (m Model) delteTask(id int64) error {
	err := m.options.Client.DeleteTask(id)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return err
}
