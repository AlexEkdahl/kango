package datastruct

import "fmt"

const BoardTableName = "board"

type Board struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Desc string `db:"description"`
}

func (b Board) String() string {
	return fmt.Sprintf("Board{ID: %d, Name: %s, Desc: %s}", b.ID, b.Name, b.Desc)
}

func (b Board) Title() string       { return b.Name }
func (b Board) Description() string { return b.Desc }
func (b Board) FilterValue() string { return b.Name }
