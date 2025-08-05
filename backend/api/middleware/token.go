package middleware

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/token"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/errno"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/response"
)

type AuthnHandler struct {
	ignore map[string]struct{}
	token  token.JWT
}

func NewAuthnHandler(token token.JWT) *AuthnHandler {
	return &AuthnHandler{token: token, ignore: make(map[string]struct{})}
}

func (h *AuthnHandler) IgnorePath(path string) *AuthnHandler {
	h.ignore[path] = struct{}{}
	return h
}

func (h *AuthnHandler) JWTAuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := h.ignore[c.Request.URL.Path]; ok {
			c.Next()
			return
		}

		access, err := h.token.GetAccessToken(c)
		if err != nil {
			response.Abort(c, errno.ErrUnauthorized)
			return
		}

		if claims, err := h.token.ParseToken(access); err == nil {
			c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "user_id", claims.UID))
			c.Next()
			return
		}

		refresh, err := c.Cookie("robot_refresh")
		if err != nil {
			response.Abort(c, errno.ErrUnauthorized)
			return
		}
		tokens, uid, err := h.token.TryRefresh(refresh, c.Request.UserAgent())
		if err != nil {
			response.Abort(c, errno.ErrUnauthorized)
			return
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "user_id", uid))

		//util.SetAuthorization(c, tokens)

		c.Next()
	}
}
