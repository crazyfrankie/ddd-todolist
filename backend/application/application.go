package application

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/application/base/appinfra"
	"github.com/crazyfrankie/ddd-todolist/backend/application/task"
	"github.com/crazyfrankie/ddd-todolist/backend/application/user"
)

type UserService = user.UserApplicationService
type TaskService = task.TaskApplicationService

type Services struct {
	infra   *appinfra.AppDependencies
	UserSvc *user.UserApplicationService
	TaskSvc *task.TaskApplicationService
}

func Init(ctx context.Context) (*Services, error) {
	infra, err := appinfra.Init(ctx)
	if err != nil {
		return nil, err
	}

	services, err := initServices(ctx, infra)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func initServices(ctx context.Context, infra *appinfra.AppDependencies) (*Services, error) {
	userSvc := user.InitService(ctx, infra.DB, infra.Storage, infra.IDGenSVC, infra.JWTGen)
	taskSvc := task.InitService(ctx, infra.DB, infra.IDGenSVC)

	return &Services{
		infra:   infra,
		UserSvc: userSvc,
		TaskSvc: taskSvc,
	}, nil
}
