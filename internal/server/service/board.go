package service

import (
	"github.com/AlexEkdahl/kango/internal/datastruct"
	"github.com/AlexEkdahl/kango/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BoardService interface {
	GetAllBoards() (*[]datastruct.Board, error)
	CreateBoard(*datastruct.Board) (*int64, error)
}

type boardService struct {
	dao repository.DAO
}

func NewBoardService(dao repository.DAO) BoardService {
	return &boardService{dao: dao}
}

func (u *boardService) GetAllBoards() (*[]datastruct.Board, error) {
	boards, err := u.dao.NewBoardQuery().GetAllBoards()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "requested user doesn't exist: %v", err)
	}

	return boards, nil
}

func (u *boardService) CreateBoard(b *datastruct.Board) (*int64, error) {
	id, err := u.dao.NewBoardMutation().CreateBoard(*b)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "requested user doesn't exist: %v", err)
	}

	return id, nil
}
