package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/internal/dal"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/internal/dal/model"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTaskList(ctx context.Context, userID int64) ([]*model.Task, error)
	GetTaskByID(ctx context.Context, taskID int64) (*model.Task, error)
	UpdateTask(ctx context.Context, taskID int64, updates map[string]any) error
	DeleteTask(ctx context.Context, taskID int64) error
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return dal.NewTaskDAO(db)
}
