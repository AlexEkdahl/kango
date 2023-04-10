package customlist

import (
	"github.com/AlexEkdahl/kango/internal/cli/keymap"
	"github.com/AlexEkdahl/kango/internal/cli/theme"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func New(width, height int) list.Model {
	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		BorderForeground(theme.AmberColour).
		Foreground(theme.AmberColour).
		Bold(true)

	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(theme.PaleAmberColour).
		BorderForeground(theme.AmberColour)

	delegate.Styles.DimmedDesc = delegate.Styles.DimmedDesc.
		Foreground(theme.FeintColour)

	delegate.Styles.NormalTitle = delegate.Styles.NormalTitle.
		Foreground(theme.FeintColour)

	delegate.Styles.FilterMatch = lipgloss.NewStyle().
		Underline(true).
		Bold(true)

	filteredList := list.New([]list.Item{}, delegate, width, height)

	// Override the colours within the existing styles
	filteredList.Styles.FilterPrompt = filteredList.Styles.FilterPrompt.
		Foreground(theme.HighlightColour)

	filteredList.Styles.FilterCursor = filteredList.Styles.FilterCursor.
		Foreground(theme.HighlightColour)

	filteredList.Styles.StatusBarFilterCount = filteredList.Styles.StatusBarFilterCount.
		Foreground(theme.FeintColour)

	filteredList.SetShowStatusBar(false)
	filteredList.SetShowTitle(false)
	filteredList.SetShowHelp(false)
	filteredList.DisableQuitKeybindings()

	// Override key bindings to force expected behaviour
	filteredList.KeyMap.GoToEnd.SetEnabled(false)
	filteredList.KeyMap.GoToStart.SetEnabled(false)

	filteredList.KeyMap.CursorUp = keymap.Up
	filteredList.KeyMap.CursorDown = keymap.Down
	filteredList.KeyMap.NextPage = keymap.Right
	filteredList.KeyMap.PrevPage = keymap.Left

	return filteredList
}

func NewBoardList(width, height int) list.Model {
	delegate := NewUnstyledDeligate()

	filteredList := list.New([]list.Item{}, delegate, width, height)

	// Override the colours within the existing styles
	filteredList.Styles.FilterPrompt = filteredList.Styles.FilterPrompt.
		Foreground(theme.HighlightColour)

	filteredList.Styles.FilterCursor = filteredList.Styles.FilterCursor.
		Foreground(theme.HighlightColour)

	filteredList.Styles.StatusBarFilterCount = filteredList.Styles.StatusBarFilterCount.
		Foreground(theme.FeintColour)

	filteredList.Styles.Title = filteredList.Styles.Title.
		Background(theme.GreenColour).
		Foreground(theme.PrimaryColour)

	filteredList.SetShowStatusBar(false)
	filteredList.SetShowTitle(true)
	filteredList.SetShowHelp(false)
	filteredList.DisableQuitKeybindings()

	// Override key bindings to force expected behaviour
	filteredList.KeyMap.GoToEnd.SetEnabled(false)
	filteredList.KeyMap.GoToStart.SetEnabled(false)

	filteredList.KeyMap.CursorUp = keymap.Up
	filteredList.KeyMap.CursorDown = keymap.Down
	filteredList.KeyMap.NextPage = keymap.Right
	filteredList.KeyMap.PrevPage = keymap.Left

	return filteredList
}

func NewUnstyledDeligate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(theme.FeintColour).
		UnsetBorderForeground().UnsetBorderLeftForeground()

	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(theme.FeintColour).
		UnsetBorderForeground().UnsetBorderLeftForeground()

	delegate.Styles.DimmedDesc = delegate.Styles.DimmedDesc.
		Foreground(theme.FeintColour).
		UnsetBorderForeground().UnsetBorderLeftForeground()

	delegate.Styles.NormalTitle = delegate.Styles.NormalTitle.
		Foreground(theme.FeintColour).
		UnsetBorderForeground().UnsetBorderLeftForeground()

	delegate.Styles.FilterMatch = lipgloss.NewStyle().
		Underline(true).
		Bold(true)

	return delegate
}

func NewStyledDeligate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		BorderForeground(theme.AmberColour).
		Foreground(theme.AmberColour).
		Bold(true)

	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(theme.PaleAmberColour).
		BorderForeground(theme.AmberColour)

	delegate.Styles.DimmedDesc = delegate.Styles.DimmedDesc.
		Foreground(theme.FeintColour)

	delegate.Styles.NormalTitle = delegate.Styles.NormalTitle.
		Foreground(theme.FeintColour)

	delegate.Styles.FilterMatch = lipgloss.NewStyle().
		Underline(true).
		Bold(true)

	return delegate
}
