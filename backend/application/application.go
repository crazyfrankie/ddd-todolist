package application

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/application/base/appinfra"
	"github.com/crazyfrankie/ddd-todolist/backend/application/task"
	"github.com/crazyfrankie/ddd-todolist/backend/application/user"
)

type services struct {
	infra   *appinfra.AppDependencies
	userSvc *user.UserApplicationService
	taskSvc *task.TaskApplicationService
}

func Init(ctx context.Context) error {
	infra, err := appinfra.Init(ctx)
	if err != nil {
		return err
	}

	_, err = initServices(ctx, infra)
	if err != nil {
		return nil
	}

	return nil
}

func initServices(ctx context.Context, infra *appinfra.AppDependencies) (*services, error) {
	userSvc := user.InitService(ctx, infra.DB, infra.Storage, infra.IDGenSVC)
	taskSvc := task.InitService(ctx, infra.DB, infra.IDGenSVC)

	return &services{
		infra:   infra,
		userSvc: userSvc,
		taskSvc: taskSvc,
	}, nil
}
