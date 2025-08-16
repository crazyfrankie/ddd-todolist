package token

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GetAccessToken(c *gin.Context) (string, error)
	GenerateToken(uid int64, ua string) ([]string, error)
	ParseToken(token string) (*jwt.MapClaims, error)
	TryRefresh(refresh string, ua string) ([]string, *jwt.MapClaims, error)
	CleanToken(ctx context.Context, uid int64, ua string) error
}
