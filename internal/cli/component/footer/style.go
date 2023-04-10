package footer

import (
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Border        lipgloss.Style
	Ellipsis      lipgloss.Style
	HelpText      lipgloss.Style
	HelpFeintText lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}

	s.Border = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(theme.BorderColour)

	s.Ellipsis = theme.FeintTextStyle.Copy()
	s.HelpText = theme.PurpleTextStyle.Copy()
	s.HelpFeintText = theme.VeryFeintTextStyle.Copy()
	return s
}
