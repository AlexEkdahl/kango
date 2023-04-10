package app

import (
	"context"

	desc "github.com/AlexEkdahl/kango/pkg/contract"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *KangoBoardServer) GetAllBoards(ctx context.Context, req *emptypb.Empty) (*desc.BoardsResponse, error) {
	boards, err := m.boardService.GetAllBoards()
	if err != nil {
		return nil, err
	}

	dtBoards := make([]*desc.Board, len(*boards))
	for i, t := range *boards {
		dtBoards[i] = &desc.Board{
			Id:          t.ID,
			Name:        t.Name,
			Description: t.Desc,
		}
	}

	return &desc.BoardsResponse{
		Boards: dtBoards,
	}, nil
}
