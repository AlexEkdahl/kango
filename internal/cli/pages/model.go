package pages

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Model interface {
	tea.Model
	help.KeyMap

	Resize(width, height int) Model
	Width() int
	Height() int
}
