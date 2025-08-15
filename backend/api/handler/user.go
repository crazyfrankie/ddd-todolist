package handler

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/api/model/user"
	"github.com/crazyfrankie/ddd-todolist/backend/application"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/logs"
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
		userGroup.GET("profile", h.GetUserInfo())
		userGroup.PUT("avatar", h.UpdateUserAvatar())
		userGroup.PUT("profile", h.UpdateUserProfile())
		userGroup.POST("reset-password", h.ResetPassword())
	}
}

// UserRegister user register
// @router /api/user/register [POST]
func (h *UserHandler) UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user.EmailRegisterRequest
		if err := c.ShouldBind(&req); err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		userInfo, tokens, err := h.svc.UserRegister(c.Request.Context(), c.Request.UserAgent(), &req)
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.Header("x-access-token", tokens[0])
		c.SetCookie("todolist_refresh", tokens[1], int(time.Hour*24), "/", "", false, true)

		data(c, userInfo)
	}
}

// UserLogin user login
// @router /api/user/login [POST]
func (h *UserHandler) UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user.EmailLoginRequest
		if err := c.ShouldBind(&req); err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		userInfo, tokens, err := h.svc.UserLogin(c.Request.Context(), c.Request.UserAgent(), &req)
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.Header("x-access-token", tokens[0])
		c.SetCookie("todolist_refresh", tokens[1], int(time.Hour*24), "/", "", false, true)

		data(c, userInfo)
	}
}

// UserLogout user logout
// @router /api/user/logout [GET]
func (h *UserHandler) UserLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h.svc.UserLogout(c.Request.Context(), c.Request.UserAgent())
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		c.SetCookie("todolist_refresh", "", int(time.Hour*24), "/", "", false, true)

		success(c)
	}
}

// GetUserInfo returns user info
// @router /api/user/profile [GET]
func (h *UserHandler) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h.svc.GetUserInfo(c.Request.Context())
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		data(c, resp)
	}
}

// UpdateUserAvatar update user avatar
// @router /api/user/avatar [POST]
func (h *UserHandler) UpdateUserAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user.UpdateAvatarRequest

		file, err := c.FormFile("avatar")
		if err != nil {
			logs.CtxErrorf(c.Request.Context(), "Get Avatar Fail failed, err=%v", err)
			invalidParamRequestResponse(c, "missing avatar file")
			return
		}

		// Check file type
		if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
			invalidParamRequestResponse(c, "invalid file type, only image allowed")
			return
		}

		// Read file content
		src, err := file.Open()
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}
		defer src.Close()

		fileContent, err := io.ReadAll(src)
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		req.Avatar = fileContent
		mimeType := file.Header.Get("Content-Type")

		url, err := h.svc.UpdateUserAvatar(c.Request.Context(), mimeType, &req)
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		data(c, gin.H{"web_uri": url})
	}
}

// UpdateUserProfile update user profile
// @router /api/user/profile [POST]
func (h *UserHandler) UpdateUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user.UpdateProfileRequest
		if err := c.ShouldBind(&req); err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		err := h.svc.UpdateUserProfile(c.Request.Context(), &req)
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		success(c)
	}
}

// ResetPassword reset user password
// @router /api/user/reset-password [POST]
func (h *UserHandler) ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req user.ResetUserPassword
		if err := c.ShouldBind(&req); err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		err := h.svc.ResetUserPassword(c.Request.Context(), &req)
		if err != nil {
			invalidParamRequestResponse(c, err.Error())
			return
		}

		success(c)
	}
}
