package app

import (
	"context"

	"github.com/AlexEkdahl/kango/internal/datastruct"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
)

func (m *MicroserviceServer) CreateTask(ctx context.Context, req *desc.Task) (*desc.TaskID, error) {
	id, err := m.taskService.CreateTask(&datastruct.Task{
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
