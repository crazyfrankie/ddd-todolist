package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/pkg/ctxcache"
)

func CtxCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(ctxcache.Init(c.Request.Context()))
		
		c.Next()
	}
}
