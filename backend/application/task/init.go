package task

import (
	"context"

	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/repository"
	"github.com/crazyfrankie/ddd-todolist/backend/domain/task/service"
	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/idgen"
)

func InitService(ctx context.Context, db *gorm.DB, idgen idgen.IDGenerator) *TaskApplicationService {
	TaskAppliCationSVC.DomainSVC = service.NewTaskDomain(ctx, &service.Components{
		IDGen:    idgen,
		TaskRepo: repository.NewTaskRepository(db),
	})

	return TaskAppliCationSVC
}
