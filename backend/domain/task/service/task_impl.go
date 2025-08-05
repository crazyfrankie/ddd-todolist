package service

import (
	"context"
	"fmt"
	"time"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/entity"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/internal/dal/model"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/repository"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/idgen"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/ptr"
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

func (t *taskImpl) CreateTask(ctx context.Context, req *CreateTaskRequest) (task *entity.Task, err error) {
	taskID, err := t.IDGen.GenID(ctx)
	if err != nil {
		return nil, fmt.Errorf("generate id error: %v", err)
	}

	now := time.Now().UnixMilli()
	newTask := &model.Task{
		ID:        taskID,
		Content:   req.Content,
		UserID:    req.UserID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if req.Date != nil {
		newTask.DueTime = ptr.From(req.Date)
	}
	if req.Priority != nil {
		newTask.Priority = ptr.From(req.Priority)
	}

	err = t.TaskRepo.CreateTask(ctx, newTask)
	if err != nil {
		return nil, err
	}

	return taskPo2Do(newTask), nil
}

func (t *taskImpl) GetTaskList(ctx context.Context, userID int64) (tasks []*entity.Task, err error) {
	taskModels, err := t.TaskRepo.GetTaskList(ctx, userID)
	if err != nil {
		return nil, err
	}

	tasks = make([]*entity.Task, 0, len(taskModels))
	for _, task := range taskModels {
		tasks = append(tasks, taskPo2Do(task))
	}

	return tasks, nil
}

func (t *taskImpl) GetTaskByID(ctx context.Context, taskID int64) (task *entity.Task, err error) {
	taskModel, err := t.TaskRepo.GetTaskByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return taskPo2Do(taskModel), nil
}

func (t *taskImpl) UpdateTask(ctx context.Context, req *UpdateTaskRequest) error {
	updates := make(map[string]any)

	if req.IsCompleted != nil {
		updates["is_completed"] = ptr.From(req.IsCompleted)
	}
	if req.Date != nil {
		updates["due_time"] = ptr.From(req.Date)
	}
	if req.Content != nil {
		updates["content"] = ptr.From(req.Content)
	}
	if req.Priority != nil {
		updates["priority"] = ptr.From(req.Priority)
	}

	err := t.TaskRepo.UpdateTask(ctx, req.TaskID, updates)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskImpl) DeleteTask(ctx context.Context, taskID int64) error {
	err := t.TaskRepo.DeleteTask(ctx, taskID)
	if err != nil {
		return err
	}

	return nil
}

func taskPo2Do(model *model.Task) *entity.Task {
	return &entity.Task{
		ID:       model.ID,
		Content:  model.Content,
		Priority: model.Priority,
		Date:     time.UnixMilli(model.DueTime).Format(time.RFC3339),
		TaskTyp:  determineTaskStatus(model),
	}
}

func determineTaskStatus(modelTask *model.Task) entity.TaskStatus {
	now := time.Now()

	if modelTask.IsCompleted {
		return entity.TaskCompleted
	}

	// 如果没有设置截止时间，默认为待完成
	if modelTask.DueTime == 0 {
		return entity.TaskToBeDone
	}

	dueTime := time.UnixMilli(modelTask.DueTime)

	// 判断是否过期（截止时间早于当前时间）
	if dueTime.Before(now) {
		return entity.TaskOverDue
	}

	// 判断是否是当天任务
	if isSameDay(dueTime, now) {
		return entity.TaskToBeDone
	}

	// 否则是已安排的任务
	return entity.TaskSchedule
}

func isSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
