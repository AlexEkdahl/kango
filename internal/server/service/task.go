package service

import (
	"github.com/AlexEkdahl/kango/internal/datastruct"
	"github.com/AlexEkdahl/kango/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskService interface {
	GetAllTasks() (*[]datastruct.Task, error)
	CreateTask(*datastruct.Task) (*int64, error)
	UpdateTask(*datastruct.Task) (*datastruct.Task, error)
	DeleteTask(int64) error
}

type taskService struct {
	dao repository.DAO
}

func NewTaskService(dao repository.DAO) TaskService {
	return &taskService{dao: dao}
}

func (u *taskService) GetAllTasks() (*[]datastruct.Task, error) {
	tasks, err := u.dao.NewTaskQuery().GetAllTasks()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "requested user doesn't exist: %v", err)
	}

	return tasks, nil
}

func (u *taskService) CreateTask(t *datastruct.Task) (*int64, error) {
	id, err := u.dao.NewTaskMutation().CreateTask(*t)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "requested user doesn't exist: %v", err)
	}

	return id, nil
}

func (u *taskService) UpdateTask(t *datastruct.Task) (*datastruct.Task, error) {
	task, err := u.dao.NewTaskMutation().UpdateTask(*t)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "requested user doesn't exist: %v", err)
	}

	return task, nil
}

func (u *taskService) DeleteTask(id int64) error {
	err := u.dao.NewTaskMutation().DeleteTask(id)
	if err != nil {
		return status.Errorf(codes.NotFound, "requested user doesn't exist: %v", err)
	}

	return nil
}
