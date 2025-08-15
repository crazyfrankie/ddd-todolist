package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
