package app

import (
	"context"

	desc "github.com/AlexEkdahl/kango/pkg/contract"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *KangoBoardServer) GetAllTasks(ctx context.Context, req *emptypb.Empty) (*desc.TasksResponse, error) {
	tasks, err := m.taskService.GetAllTasks()
	if err != nil {
		return nil, err
	}

	pt := make([]*desc.Task, len(*tasks))
	for i, t := range *tasks {
		pt[i] = &desc.Task{
			Id:          t.ID,
			Subject:     t.Subject,
			Description: t.Desc,
			Status:      desc.Status(t.Status),
		}
	}

	return &desc.TasksResponse{
		Tasks: pt,
	}, nil
}
