package create

import (
	"github.com/AlexEkdahl/kango/internal/cli/keymap"
	"github.com/AlexEkdahl/kango/internal/cli/message"
	"github.com/AlexEkdahl/kango/internal/cli/pages"

	// "github.com/AlexEkdahl/kango/internal/cli/theme"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	title    textinput.Model
	desc     textarea.Model
	viewport viewport.Model
	styles   *Styles
}

func New() Model {
	return Model{
		viewport: viewport.New(0, 0),
		styles:   DefaultStyles(),
		title:    textinput.New(),
		desc:     textarea.New(),
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
	case message.ShowCreateModal:
		m.title.Focus()
		m.desc.Blur()
		m.desc.ShowLineNumbers = false
		m.title.SetValue("")
		m.desc.SetValue("")
		cmds = append(cmds, textarea.Blink)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Enter):
			if m.title.Focused() {
				m.title.Blur()
				m.desc.Focus()
				return m, textarea.Blink
			}
		case key.Matches(msg, keymap.Create):
			cmds = append(cmds, func() tea.Msg {
				return message.CreateTask{
					Status: 0,
					Title:  m.title.Value(),
					Desc:   m.desc.Value(),
					Board:  0,
				}
			})
		case key.Matches(msg, keymap.Escape):
			m.title.SetValue("")
			m.desc.SetValue("")
			cmds = append(cmds, message.CloseDialogCmd)

		case key.Matches(msg, keymap.Quit):
			cmds = append(cmds, tea.Quit)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)

	cmds = append(cmds, cmd)

	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		m.desc, cmd = m.desc.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) ShortHelp() []key.Binding {
	kb := make([]key.Binding, 0)
	kb = append(kb, keymap.Enter, keymap.Create, keymap.Escape)

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

func (m Model) View() string {
	test := lipgloss.JoinVertical(lipgloss.Left, m.title.View(), m.desc.View())

	ui := lipgloss.JoinVertical(lipgloss.Left, test)

	dialog := lipgloss.Place(m.viewport.Width, m.viewport.Height,
		lipgloss.Center, lipgloss.Center,
		m.styles.DialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(m.styles.Subtle),
	)

	m.viewport.SetContent(dialog)

	return m.viewport.View()
}
