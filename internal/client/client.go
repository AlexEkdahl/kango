package client

import (
	"context"
	"fmt"

	"github.com/AlexEkdahl/kango/config"
	"github.com/AlexEkdahl/kango/internal/datastruct"
	"github.com/AlexEkdahl/kango/internal/repository"
	"github.com/AlexEkdahl/kango/internal/server/service"
	"github.com/AlexEkdahl/kango/pkg/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	GetAllTasks() (*[]datastruct.Task, error)
	// CreateTask(datastruct.Task) (*int64, error)
	// UpdateTask(datastruct.Task) (*datastruct.Task, error)
	// DeleteTask(int64) error
}

type RemoteHost struct {
	api contract.KanbanClient
}

func (r *RemoteHost) GetAllTasks() (*[]datastruct.Task, error) {
	res, err := r.api.GetAllTasks(context.TODO(), &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("Error getting tasks from server: %e", err)
	}

	tasks := []datastruct.Task{}
	for _, task := range res.GetTasks() {
		t := datastruct.Task{
			ID:      task.GetId(),
			Status:  datastruct.Status(task.Status),
			Subject: task.Subject,
			Desc:    task.Description,
		}
		tasks = append(tasks, t)
	}

	return &tasks, nil
}

func New(c config.Config) (Client, error) {
	if c.Host != "" {
		connStr := fmt.Sprintf("%s:%d", c.Host, c.Port)

		conn, err := grpc.Dial(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}

		// Create a new KanbanClient instance.
		c := contract.NewKanbanClient(conn)
		return &RemoteHost{
			api: c,
		}, nil
	}

	err := repository.NewLocalDB(c)
	if err != nil {
		return nil, err
	}
	d := repository.NewDAO()
	t := service.NewTaskService(d)

	return t, nil
}
