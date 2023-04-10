package footer

import (
	"strings"

	"github.com/AlexEkdahl/kango/internal/cli/component"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	help   help.Model
	keymap help.KeyMap
	width  int
	Styles *Styles
}

func New(keymap help.KeyMap) Model {
	styles := DefaultStyles()

	help := help.New()
	help.Styles.ShortSeparator = styles.Ellipsis
	help.Styles.ShortKey = styles.HelpText
	help.Styles.ShortDesc = styles.HelpFeintText

	return Model{
		help:   help,
		keymap: keymap,
		Styles: styles,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	panel := lipgloss.JoinVertical(lipgloss.Top, m.help.View(m.keymap))

	b.WriteString(m.Styles.Border.
		Width(m.width).
		Render(panel))

	return b.String()
}

func (m Model) Resize(width, height int) component.Model {
	m.width = width
	return m
}

func (m Model) Width() int {
	return m.width
}

func (m Model) Height() int {
	return lipgloss.Height(m.View())
}

func (m Model) SetKeyMap(keymap help.KeyMap) Model {
	m.keymap = keymap
	return m
}
