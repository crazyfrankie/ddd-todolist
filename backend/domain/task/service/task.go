package service

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/entity"
)

type CreateTaskRequest struct {
	Content  string
	Date     int64
	UserID   int64
	Priority string
}

type UpdateTaskRequest struct {
	Content  string
	Date     int64
	Priority string
}

type Task interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest) error
	GetTaskList(ctx context.Context, userID int64) ([]*entity.Task, error)
	GetTaskByID(ctx context.Context, taskID int64) (*entity.Task, error)
	UpdateTaskStatus(ctx context.Context, req *UpdateTaskRequest) error
	DeleteTask(ctx context.Context, taskID int64) error
}
