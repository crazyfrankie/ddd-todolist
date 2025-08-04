package task

import (
	task "github.com/crazyfrankie/ddd-todolist/backend/domain/task/service"
)

var TaskAppliCationSVC = &TaskApplicationService{}

type TaskApplicationService struct {
	DomainSVC task.Task
}
