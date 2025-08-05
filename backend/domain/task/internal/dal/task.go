package dal

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/internal/dal/model"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/internal/dal/query"
)

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{
		query: query.Use(db),
	}
}

type TaskDAO struct {
	query *query.Query
}

func (t *TaskDAO) CreateTask(ctx context.Context, task *model.Task) error {
	return t.query.WithContext(ctx).Task.Create(task)
}

func (t *TaskDAO) GetTaskList(ctx context.Context, userID int64) ([]*model.Task, error) {
	tasks, err := t.query.WithContext(ctx).Task.Where(t.query.Task.UserID.Eq(userID)).Find()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskDAO) GetTaskByID(ctx context.Context, taskID int64) (*model.Task, error) {
	task, err := t.query.WithContext(ctx).Task.Where(t.query.Task.ID.Eq(taskID)).First()
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskDAO) UpdateTask(ctx context.Context, taskID int64, updates map[string]any) error {
	if _, ok := updates["updated_at"]; !ok {
		updates["updated_at"] = time.Now().UnixMilli()
	}

	_, err := t.query.WithContext(ctx).Task.Where(t.query.Task.ID.Eq(taskID)).Updates(updates)
	return err
}

func (t *TaskDAO) DeleteTask(ctx context.Context, taskID int64) error {
	return t.query.Transaction(func(tx *query.Query) error {
		task, err := tx.WithContext(ctx).Task.Where(t.query.Task.ID.Eq(taskID)).First()
		if err != nil {
			return err
		}

		_, err = tx.WithContext(ctx).Task.Delete(task)
		if err != nil {
			return err
		}

		return nil
	})
}
