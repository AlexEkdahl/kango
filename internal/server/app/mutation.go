package app

import (
	"context"

	"github.com/AlexEkdahl/kango/internal/datastruct"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *KangoBoardServer) CreateTask(ctx context.Context, req *desc.Task) (*desc.TaskID, error) {
	id, err := m.TaskService.CreateTask(&datastruct.Task{
		Status:  datastruct.Status(req.GetStatus()),
		Subject: req.Subject,
		Desc:    req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	return &desc.TaskID{
		Id: *id,
	}, nil
}

func (m *KangoBoardServer) UpdateTask(ctx context.Context, req *desc.Task) (*desc.Task, error) {
	task, err := m.TaskService.UpdateTask(&datastruct.Task{
		ID:      req.GetId(),
		Status:  datastruct.Status(req.GetStatus()),
		Subject: req.Subject,
		Desc:    req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	return &desc.Task{
		Id:          task.ID,
		Subject:     task.Subject,
		Description: task.Desc,
		Status:      desc.Status(task.Status),
	}, nil
}

func (m *KangoBoardServer) DeleteTask(ctx context.Context, req *desc.TaskID) (*emptypb.Empty, error) {
	err := m.TaskService.DeleteTask(req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
