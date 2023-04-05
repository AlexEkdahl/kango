package repository

import (
	"fmt"

	"github.com/AlexEkdahl/kango/internal/datastruct"
	"github.com/Masterminds/squirrel"
)

type TaskQuery interface {
	GetAllTasks() (*[]datastruct.Task, error)
}
type TaskMutation interface {
	CreateTask(datastruct.Task) (*int64, error)
	UpdateTask(datastruct.Task) (*datastruct.Task, error)
	DeleteTask(int64) error
}

type (
	taskQuery    struct{}
	taskMutation struct{}
)

func (t *taskQuery) GetAllTasks() (*[]datastruct.Task, error) {
	qb := pgQb().
		Select("*").
		From(datastruct.TaskTableName)

	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	var tasks []datastruct.Task
	var task datastruct.Task

	for rows.Next() {
		err = rows.Scan(&task.ID, &task.Status, &task.Subject, &task.Desc)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func (t *taskMutation) CreateTask(dto datastruct.Task) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.TaskTableName).
		Columns("subject", "description", "status").
		Values(dto.Subject, dto.Desc, dto.Status).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (t *taskMutation) UpdateTask(dto datastruct.Task) (*datastruct.Task, error) {
	qb := pgQb().
		Update(datastruct.TaskTableName).
		SetMap(map[string]interface{}{
			"subject":     dto.Subject,
			"description": dto.Desc,
			"status":      dto.Status,
		}).
		Where(squirrel.Eq{"id": dto.ID}).Suffix("RETURNING id, description, status, subject")

	var updatedTask datastruct.Task
	err := qb.QueryRow().Scan(&updatedTask.ID,
		&updatedTask.Desc, &updatedTask.Status, &updatedTask.Subject)
	if err != nil {
		return nil, fmt.Errorf("cannot update the course %v", err)
	}

	return &updatedTask, nil
}

func (t *taskMutation) DeleteTask(id int64) error {
	qb := pgQb().
		Delete(datastruct.TaskTableName).
		From(datastruct.TaskTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
