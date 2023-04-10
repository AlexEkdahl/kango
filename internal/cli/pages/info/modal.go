package info

import (
	"github.com/AlexEkdahl/kango/internal/cli/keymap"
	"github.com/AlexEkdahl/kango/internal/cli/message"
	"github.com/AlexEkdahl/kango/internal/cli/pages"
	"github.com/AlexEkdahl/kango/internal/cli/theme"

	// "github.com/AlexEkdahl/kango/internal/cli/theme"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	title    string
	desc     string
	viewport viewport.Model
	styles   *Styles
}

func New() Model {
	return Model{
		viewport: viewport.New(0, 0),
		styles:   DefaultStyles(),
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
	case message.ShowItem:
		m.title = msg.Title
		m.desc = msg.Desc
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Escape):
			fallthrough
		case key.Matches(msg, keymap.Close):
			cmds = append(cmds, message.CloseDialogCmd)
		case key.Matches(msg, keymap.Quit):
			cmds = append(cmds, tea.Quit)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) ShortHelp() []key.Binding {
	kb := make([]key.Binding, 0)
	kb = append(kb, keymap.Close)

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
	taskStyle := lipgloss.NewStyle().Width(m.viewport.Width / 2).Height(m.viewport.Height / 5).Align(lipgloss.Center)

	title := taskStyle.Foreground(theme.RedColour).Render(m.title)
	desc := taskStyle.Foreground(theme.LightRedColour).Render(m.desc)

	ui := lipgloss.JoinVertical(lipgloss.Left, title, desc)

	dialog := lipgloss.Place(m.viewport.Width, m.viewport.Height,
		lipgloss.Center, lipgloss.Center,
		m.styles.DialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(m.styles.Subtle),
	)

	m.viewport.SetContent(dialog)

	return m.viewport.View()
}
