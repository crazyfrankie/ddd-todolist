package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/application"
)

type TaskHandler struct {
	svc *application.TaskService
}

func NewTaskHandler(svc *application.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (h *TaskHandler) RegisterRoute(r *gin.RouterGroup) {

}
