package repository

import (
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/internal/dal"
	"gorm.io/gorm"
)

type TaskRepository interface {
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return dal.NewTaskDAO(db)
}
