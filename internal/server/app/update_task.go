package app

import (
	"context"

	"github.com/AlexEkdahl/kango/internal/datastruct"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
)

func (m *MicroserviceServer) UpdateTask(ctx context.Context, req *desc.Task) (*desc.Task, error) {
	task, err := m.taskService.UpdateTask(&datastruct.Task{
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
