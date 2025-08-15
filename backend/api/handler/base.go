package handler

import (
	"github.com/crazyfrankie/ddd-todolist/backend/api/httputil"
	"github.com/gin-gonic/gin"
)

func invalidParamRequestResponse(c *gin.Context, errMsg string) {
	httputil.BadRequest(c, errMsg)
}

func internalServerErrorResponse(c *gin.Context, err error) {
	httputil.InternalError(c, err)
}

func success(c *gin.Context) {
	httputil.Success(c, nil)
}

func data(c *gin.Context, data any) {
	httputil.Success(c, data)
}
