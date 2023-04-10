package datastruct

import "fmt"

const TaskTableName = "task"

type Task struct {
	ID      int64  `db:"id"`
	Status  Status `db:"status"`
	Subject string `db:"subject"`
	Desc    string `db:"description"`
}

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (t Task) String() string {
	return fmt.Sprintf("Task{ID: %d, Status: %d, Subject: %q, Desc: %q}", t.ID, t.Status, t.Subject, t.Desc)
}

func (i Task) Title() string       { return i.Subject }
func (i Task) Description() string { return i.Desc }
func (i Task) FilterValue() string { return i.Subject }
