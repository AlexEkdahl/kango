package keymap

import "github.com/charmbracelet/bubbles/key"

var (
	Up = key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("k", "up"),
	)

	Down = key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("j", "down"),
	)

	UpDown = key.NewBinding(
		key.WithKeys("up", "k", "down", "j"),
		key.WithHelp("j/k", "up/down"),
	)

	Left = key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("h", "prev"),
	)

	Right = key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("l", "next"),
	)

	LeftRight = key.NewBinding(
		key.WithKeys("left", "h", "right", "l"),
		key.WithHelp("h/l", "prev/next"),
	)

	Enter = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↲", "select"),
	)

	Create = key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl+n", "send"),
	)

	ForwardSlash = key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	)

	Escape = key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	)

	Delete = key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	)
	Quit = key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c", "quit"),
	)

	Close = key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "close"),
	)

	New = key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new"),
	)
)
