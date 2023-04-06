package app

import (
	"context"

	desc "github.com/AlexEkdahl/kango/pkg/contract"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *KangoBoardServer) DeleteTask(ctx context.Context, req *desc.TaskID) (*emptypb.Empty, error) {
	err := m.taskService.DeleteTask(req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
