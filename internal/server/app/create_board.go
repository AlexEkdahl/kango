package app

import (
	"context"

	"github.com/AlexEkdahl/kango/internal/datastruct"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
)

func (m *KangoBoardServer) CreateBoard(ctx context.Context, req *desc.Board) (*desc.ResourceID, error) {
	id, err := m.boardService.CreateBoard(&datastruct.Board{
		Name: req.Name,
		Desc: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &desc.ResourceID{
		Id: *id,
	}, nil
}
