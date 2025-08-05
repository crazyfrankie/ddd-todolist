package task

import (
	"context"

	model "github.com/crazyfrankie/ddd-todolist/backend/api/model/task"
	"github.com/crazyfrankie/ddd-todolist/backend/application/base/ctxutil"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/entity"
	task "github.com/crazyfrankie/ddd-todolist/backend/domain/task/service"
)

type TaskApplicationService struct {
	DomainSVC task.Task
}

func (t *TaskApplicationService) AddTask(ctx context.Context, req *model.CreateTaskRequest) (resp *model.Task, err error) {
	userID := ctxutil.MustGetUIDFromCtx(ctx)

	taskInfo, err := t.DomainSVC.CreateTask(ctx, &task.CreateTaskRequest{
		Content:  req.Content,
		Date:     req.Date,
		UserID:   userID,
		Priority: req.Priority,
	})
	if err != nil {
		return nil, err
	}

	return taskDo2To(taskInfo), nil
}

func (t *TaskApplicationService) GetTaskDetail(ctx context.Context, taskID int64) (resp *model.Task, err error) {
	taskInfo, err := t.DomainSVC.GetTaskByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return taskDo2To(taskInfo), nil
}

func (t *TaskApplicationService) GetTaskList(ctx context.Context) (resp []*model.TaskItem, err error) {
	userID := ctxutil.MustGetUIDFromCtx(ctx)

	tasks, err := t.DomainSVC.GetTaskList(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp = make([]*model.TaskItem, 0, len(tasks))
	for _, task := range tasks {
		resp = append(resp, &model.TaskItem{
			Content: task.Content,
			TaskTyp: task.TaskTyp.String(),
		})
	}

	return resp, nil
}

func (t *TaskApplicationService) UpdateTask(ctx context.Context, req *model.UpdateTaskRequest) error {
	err := t.DomainSVC.UpdateTask(ctx, &task.UpdateTaskRequest{
		TaskID:      req.TaskID,
		Content:     req.Content,
		Date:        req.Date,
		Priority:    req.Priority,
		IsCompleted: req.IsCompleted,
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskApplicationService) DeleteTask(ctx context.Context, taskID int64) error {
	err := t.DomainSVC.DeleteTask(ctx, taskID)
	if err != nil {
		return err
	}

	return nil
}

func taskDo2To(taskDo *entity.Task) *model.Task {
	return &model.Task{
		ID:      taskDo.ID,
		Content: taskDo.Content,
		Date:    taskDo.Date,
		TaskTyp: taskDo.TaskTyp.String(),
	}
}
