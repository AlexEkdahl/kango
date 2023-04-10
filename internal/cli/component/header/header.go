package header

import (
	"strings"

	"github.com/AlexEkdahl/kango/internal/cli/component"
	"github.com/AlexEkdahl/kango/internal/cli/message"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	name        string
	version     string
	description string
	width       int
	Styles      *Styles
	tempVersion string
	tempName    string
}

func New(name, version, description string) Model {
	return Model{
		name:        name,
		description: description,
		version:     version,
		Styles:      DefaultStyles(),
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

	header := lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.Styles.Name.Padding(0, 2).Render(m.name),
		m.Styles.Version.Padding(0, 2).Render(m.version),
		m.Styles.Description.Padding(0, 2).Render(m.description),
	)

	banner := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Render(header),
	)

	b.WriteString(m.Styles.Border.
		MarginBottom(1).
		Width(m.width).
		Render(banner))

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

func (m Model) UpdateName(msg message.SelectedBoard) Model {
	m.tempName = m.name
	m.tempVersion = m.version

	m.name = msg.Name
	m.version = msg.Desc

	return m
}

func (m Model) Reset() Model {
	m.name = m.tempName
	m.version = m.tempVersion

	return m
}
