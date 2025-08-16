package token

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GetAccessToken(c *gin.Context) (string, error)
	GenerateToken(uid int64, ua string) ([]string, error)
	ParseToken(token string) (*Claims, error)
	TryRefresh(refresh string, ua string) ([]string, *Claims, error)
	CleanToken(ctx context.Context, uid int64, ua string) error
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}
