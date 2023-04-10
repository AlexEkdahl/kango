package app

import (
	"context"

	"github.com/AlexEkdahl/kango/internal/datastruct"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
)

func (m *KangoBoardServer) CreateTask(ctx context.Context, req *desc.Task) (*desc.ResourceID, error) {
	id, err := m.taskService.CreateTask(&datastruct.Task{
		Status:  datastruct.Status(req.GetStatus()),
		Subject: req.Subject,
		Desc:    req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	return &desc.ResourceID{
		Id: *id,
	}, nil
}
