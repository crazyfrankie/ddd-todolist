package httputil

import (
	"errors"
	"net/http"

	"github.com/crazyfrankie/frx/errorx"
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/pkg/logs"
)

type data struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func BadRequest(c *gin.Context, errMsg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, data{Code: http.StatusBadRequest, Message: errMsg})
}

func InternalError(c *gin.Context, err error) {
	var customErr errorx.StatusError

	if errors.As(err, &customErr) && customErr.Code() != 0 {
		logs.CtxWarnf(c.Request.Context(), "[ErrorX] error:  %v %v \n", customErr.Code(), err)
		c.AbortWithStatusJSON(http.StatusOK, data{Code: customErr.Code(), Message: customErr.Msg()})
		return
	}

	logs.CtxErrorf(c.Request.Context(), "[InternalError]  error: %v \n", err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, data{Code: 500, Message: "internal server error"})
}
