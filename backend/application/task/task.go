package task

import (
	task "github.com/crazyfrankie/ddd-todolist/backend/domain/task/service"
)

type TaskApplicationService struct {
	DomainSVC task.Task
}
