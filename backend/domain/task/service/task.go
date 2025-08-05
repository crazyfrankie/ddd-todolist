package service

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/entity"
)

type CreateTaskRequest struct {
	Content  string
	UserID   int64
	Date     *int64
	Priority *string
}

type UpdateTaskRequest struct {
	TaskID      int64
	Content     *string
	Date        *int64
	Priority    *string
	IsCompleted *bool
}

type Task interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest) (*entity.Task, error)
	GetTaskList(ctx context.Context, userID int64) ([]*entity.Task, error)
	GetTaskByID(ctx context.Context, taskID int64) (*entity.Task, error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest) error
	DeleteTask(ctx context.Context, taskID int64) error
}
