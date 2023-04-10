package theme

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	PrimaryColour        = lipgloss.Color("#282c34") // One Dark Pro background
	SecondaryColour      = lipgloss.Color("#3b4048") // One Dark Pro lighter background
	BorderColour         = lipgloss.Color("#434a54") // One Dark Pro border
	FeintColour          = lipgloss.Color("#5c6370") // One Dark Pro comment
	VeryFeintColour      = lipgloss.Color("#4b526d") // One Dark Pro inactive selection background
	TextColour           = lipgloss.Color("#abb2bf") // One Dark Pro foreground
	HighlightColour      = lipgloss.Color("#528bff") // One Dark Pro blue
	HighlightFeintColour = lipgloss.Color("#61afef") // One Dark Pro cyan
	AmberColour          = lipgloss.Color("#e5c07b") // One Dark Pro yellow
	GreenColour          = lipgloss.Color("#98c379") // One Dark Pro green
	RedColour            = lipgloss.Color("#e06c75") // One Dark Pro red
	OrangeColour         = lipgloss.Color("#d19a66") // One Dark Pro orange
	LightBlueColour      = lipgloss.Color("#56b6c2") // One Dark Pro light blue
	LightGreenColour     = lipgloss.Color("#809967") // One Dark Pro light green
	LightRedColour       = lipgloss.Color("#f27c7c") // One Dark Pro light red
	PurpleColour         = lipgloss.Color("#c678dd") // One Dark Pro purple
	PaleAmberColour      = lipgloss.Color("#a68b5a") // One Dark Pro yellow
	PalePurpleColour     = lipgloss.Color("#7d7196") // One Dark Pro light purple

	AppStyle            = lipgloss.NewStyle().Margin(1)
	TextStyle           = lipgloss.NewStyle().Foreground(AmberColour)
	GreenTextStyle      = lipgloss.NewStyle().Foreground(GreenColour)
	FeintTextStyle      = lipgloss.NewStyle().Foreground(FeintColour)
	VeryFeintTextStyle  = lipgloss.NewStyle().Foreground(VeryFeintColour)
	HightlightTextStyle = lipgloss.NewStyle().Foreground(HighlightColour)

	OrangeTextStyle     = lipgloss.NewStyle().Foreground(OrangeColour)
	LightBlueTextStyle  = lipgloss.NewStyle().Foreground(LightBlueColour)
	LightGreenTextStyle = lipgloss.NewStyle().Foreground(LightGreenColour)
	LightRedTextStyle   = lipgloss.NewStyle().Foreground(LightRedColour)
	PurpleTextStyle     = lipgloss.NewStyle().Foreground(PurpleColour)
	PalePurpleTextStyle = lipgloss.NewStyle().Foreground(PalePurpleColour)
	PaleAmberTextStyle  = lipgloss.NewStyle().Foreground(PaleAmberColour)
)
