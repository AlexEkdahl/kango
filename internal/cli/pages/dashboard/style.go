package dashboard

import (
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Spinner lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}

	s.Spinner = lipgloss.NewStyle().
		Foreground(theme.HighlightColour).
		Bold(true)

	return s
}
