package service

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/entity"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/repository"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/idgen"
)

type Components struct {
	IDGen    idgen.IDGenerator
	TaskRepo repository.TaskRepository
}

type taskImpl struct {
	*Components
}

func NewTaskDomain(ctx context.Context, c *Components) Task {
	return &taskImpl{
		Components: c,
	}
}

func (t *taskImpl) CreateTask(ctx context.Context, req *CreateTaskRequest) error {
	//TODO implement me
	panic("implement me")
}

func (t *taskImpl) GetTaskList(ctx context.Context, userID int64) ([]*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskImpl) GetTaskByID(ctx context.Context, taskID int64) (*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskImpl) UpdateTaskStatus(ctx context.Context, req *UpdateTaskRequest) error {
	//TODO implement me
	panic("implement me")
}

func (t *taskImpl) DeleteTask(ctx context.Context, taskID int64) error {
	//TODO implement me
	panic("implement me")
}
