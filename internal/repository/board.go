package repository

import (
	"github.com/AlexEkdahl/kango/internal/datastruct"
)

type BoardQuery interface {
	GetAllBoards() (*[]datastruct.Board, error)
}

type BoardMutation interface {
	CreateBoard(datastruct.Board) (*int64, error)
}

type (
	boardQuery    struct{}
	boardMutation struct{}
)

func (bq *boardQuery) GetAllBoards() (*[]datastruct.Board, error) {
	qb := pgQb().
		Select("*").
		From(datastruct.BoardTableName)

	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	var boards []datastruct.Board
	var board datastruct.Board

	for rows.Next() {
		err = rows.Scan(&board.ID, &board.Name, &board.Desc)
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}

	return &boards, nil
}

func (bm *boardMutation) CreateBoard(dto datastruct.Board) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.BoardTableName).
		Columns("name", "description").
		Values(dto.Name, dto.Desc).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
