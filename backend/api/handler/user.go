package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/application"
)

type UserHandler struct {
	svc *application.UserService
}

func NewUserHandler(svc *application.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoute(r *gin.RouterGroup) {

}
