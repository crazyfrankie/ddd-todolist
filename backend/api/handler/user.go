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
	userGroup := r.Group("user")
	{
		userGroup.POST("register", h.UserRegister())
		userGroup.POST("login", h.UserLogin())
		userGroup.GET("logout", h.UserLogout())
		userGroup.PUT("avatar", h.UpdateUserAvatar())
		userGroup.PUT("profile", h.UpdateUserProfile())
		userGroup.POST("reset-password", h.ResetPassword())
	}
}

func (h *UserHandler) UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func (h *UserHandler) UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func (h *UserHandler) UserLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func (h *UserHandler) UpdateUserAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func (h *UserHandler) UpdateUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func (h *UserHandler) ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}
