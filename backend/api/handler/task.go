package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/api/model/task"
	"github.com/crazyfrankie/ddd-todolist/backend/application"
)

type TaskHandler struct {
	svc *application.TaskService
}

func NewTaskHandler(svc *application.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (h *TaskHandler) RegisterRoute(r *gin.RouterGroup) {
	taskGroup := r.Group("tasks")
	{
		taskGroup.POST("", h.AddTask())
		taskGroup.GET("/:task_id", h.GetTaskDetail())
		taskGroup.GET("", h.GetTaskList())
		taskGroup.PUT("", h.UpdateTask())
		taskGroup.DELETE("/:task_id", h.DeleteTask())
	}
}

// AddTask create a new task
// @router /api/tasks [POST]
func (h *TaskHandler) AddTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req task.CreateTaskRequest
		if err := c.ShouldBind(&req); err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		resp, err := h.svc.AddTask(c.Request.Context(), &req)
		if err != nil {
			internalServerErrorResponse(c, err)
			return
		}

		data(c, resp)
	}
}

// GetTaskDetail get a task detail
// @router /api/task/:task_id [GET]
func (h *TaskHandler) GetTaskDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskIDStr := c.Param("task_id")
		taskID, _ := strconv.ParseInt(taskIDStr, 10, 64)

		resp, err := h.svc.GetTaskDetail(c.Request.Context(), taskID)
		if err != nil {
			internalServerErrorResponse(c, err)
			return
		}

		data(c, resp)
	}
}

// GetTaskList returns tasks list
// @router /api/tasks [GET]
func (h *TaskHandler) GetTaskList() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h.svc.GetTaskList(c.Request.Context())
		if err != nil {
			internalServerErrorResponse(c, err)
			return
		}

		data(c, resp)
	}
}

// UpdateTask update task info
// @router /api/tasks [PUT]
func (h *TaskHandler) UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req task.UpdateTaskRequest
		if err := c.ShouldBind(&req); err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		err := h.svc.UpdateTask(c.Request.Context(), &req)
		if err != nil {
			internalServerErrorResponse(c, err)
			return
		}

		success(c)
	}
}

// DeleteTask delete a task
// @router /api/tasks/:task_id [DELETE]
func (h *TaskHandler) DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskIDStr := c.Param("task_id")
		taskID, _ := strconv.ParseInt(taskIDStr, 10, 64)

		err := h.svc.DeleteTask(c.Request.Context(), taskID)
		if err != nil {
			internalServerErrorResponse(c, err)
			return
		}

		success(c)
	}
}
