package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetLogID() gin.HandlerFunc {
	return func(c *gin.Context) {
		logID := uuid.New()
		ctx := context.WithValue(c.Request.Context(), "log-id", logID)

		c.Request = c.Request.WithContext(ctx)
	}
}
