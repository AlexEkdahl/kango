package kanbanboard

import (
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	ColumnStyle  lipgloss.Style
	FocusedStyle lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}

	s.ColumnStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.HiddenBorder())

	s.FocusedStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(theme.PurpleColour)

	return s
}
