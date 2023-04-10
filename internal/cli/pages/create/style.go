package create

import (
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Subtle            lipgloss.AdaptiveColor
	DialogBoxStyle    lipgloss.Style
	ButtonStyle       lipgloss.Style
	ActiveButtonStyle lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}

	s.Subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	s.DialogBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.PurpleColour).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	s.ButtonStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginTop(1)

	s.ActiveButtonStyle = s.ButtonStyle.Copy().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#F25D94")).
		MarginRight(2).
		Underline(true)

	return s
}
