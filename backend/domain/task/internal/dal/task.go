package dal

import (
	"gorm.io/gorm"

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
