package message

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ErrorMsg struct {
	Reason string
	Cause  error
}
type RefreshKeymap struct{}

type CloseDialog struct{}

type GoBack struct{}

type ShowCreateModal struct{}

type CreateTask struct {
	ID     int64
	Status int64
	Title  string
	Desc   string
	Board  int
}

type SelectedBoard struct {
	ID   int64
	Name string
	Desc string
}

type ShowItem struct {
	ID    int64
	Title string
	Desc  string
}

func RefreshKeyMapCmd() tea.Msg   { return RefreshKeymap{} }
func CloseDialogCmd() tea.Msg     { return CloseDialog{} }
func GoBackCmd() tea.Msg          { return GoBack{} }
func ShowCreateModalCmd() tea.Msg { return ShowCreateModal{} }

// func CreateTaskCmd() tea.Msg      { return CreateTask{} }
