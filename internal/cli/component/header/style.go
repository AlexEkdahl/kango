package header

import (
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Name        lipgloss.Style
	Description lipgloss.Style
	Version     lipgloss.Style
	Border      lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}
	s.Name = lipgloss.NewStyle().
		Foreground(theme.TextColour).
		Background(theme.PrimaryColour).
		Bold(true)

	s.Description = lipgloss.NewStyle().
		Foreground(theme.PalePurpleColour).
		Bold(false)

	s.Version = lipgloss.NewStyle().
		Foreground(theme.TextColour).
		Background(theme.SecondaryColour)

	s.Border = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, false).
		BorderForeground(theme.BorderColour)

	return s
}
