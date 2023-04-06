package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/AlexEkdahl/kango/internal/datastruct"
	"github.com/AlexEkdahl/kango/internal/repository"
	"github.com/AlexEkdahl/kango/internal/server/app"
	"github.com/AlexEkdahl/kango/internal/server/service"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
)

type mockDAO struct{}

func (m *mockDAO) NewTaskMutation() repository.TaskMutation {
	return &mockTaskMutation{}
}

func (m *mockDAO) NewTaskQuery() repository.TaskQuery {
	return &mockTaskQuery{}
}

type (
	mockTaskMutation struct{}
	mockTaskQuery    struct{}
)

func (m *mockTaskQuery) GetAllTasks() (*[]datastruct.Task, error) {
	tasks := []datastruct.Task{}
	return &tasks, nil
}

func (m *mockTaskMutation) CreateTask(task datastruct.Task) (*int64, error) {
	if task.Subject == "Test Task" {
		id := int64(1)
		return &id, nil
	}
	return nil, errors.New("invalid task subject")
}

func (m *mockTaskMutation) UpdateTask(task datastruct.Task) (*datastruct.Task, error) {
	if task.ID == 1 {
		return &task, nil
	}
	return nil, errors.New("invalid task ID")
}

func (m *mockTaskMutation) DeleteTask(id int64) error {
	if id == 1 {
		return nil
	}
	return errors.New("invalid task ID")
}

// Implement other methods required by the repository.DAO and repository.TaskMutation interfaces

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	req := &desc.Task{
		Status:      desc.Status_TODO,
		Subject:     "Test Task",
		Description: "This is a test task",
	}

	mockDAO := &mockDAO{}
	taskService := service.NewTaskService(mockDAO)
	kangoBoardServer := app.NewMicroservice(taskService)

	resp, err := kangoBoardServer.CreateTask(ctx, req)
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	if resp.GetId() != 1 {
		t.Errorf("Expected task ID to be 1, but got %d", resp.GetId())
	}
}

func TestDeleteTask(t *testing.T) {
	ctx := context.Background()
	req := &desc.TaskID{
		Id: 1,
	}

	mockDAO := &mockDAO{}
	taskService := service.NewTaskService(mockDAO)
	kangoBoardServer := app.NewMicroservice(taskService)

	_, err := kangoBoardServer.DeleteTask(ctx, req)
	if err != nil {
		t.Fatalf("DeleteTask failed: %v", err)
	}
}

func TestUpdateTask(t *testing.T) {
	ctx := context.Background()
	req := &desc.Task{
		Id:          1,
		Status:      desc.Status_TODO,
		Subject:     "Updated Task",
		Description: "This is an updated test task",
	}

	mockDAO := &mockDAO{}
	taskService := service.NewTaskService(mockDAO)
	kangoBoardServer := app.NewMicroservice(taskService)

	resp, err := kangoBoardServer.UpdateTask(ctx, req)
	if err != nil {
		t.Fatalf("UpdateTask failed: %v", err)
	}

	if resp.GetSubject() != "Updated Task" {
		t.Errorf("Expected task subject to be 'Updated Task', but got %s", resp.GetSubject())
	}
}
