package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/ddd-todolist/backend/infra/contract/token"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/ctxcache"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/errno"
	"github.com/crazyfrankie/ddd-todolist/backend/pkg/response"
	"github.com/crazyfrankie/ddd-todolist/backend/types/consts"
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
			ctxcache.Store(c.Request.Context(), consts.SessionDataKeyInCtx, int64((*claims)["user_id"].(float64)))
			c.Next()
			return
		}

		refresh, err := c.Cookie("todolist_refresh")
		if err != nil {
			response.Abort(c, errno.ErrUnauthorized)
			return
		}
		tokens, uid, err := h.token.TryRefresh(refresh, c.Request.UserAgent())
		if err != nil {
			response.Abort(c, errno.ErrUnauthorized)
			return
		}
		ctxcache.Store(c.Request.Context(), consts.SessionDataKeyInCtx, uid)

		c.SetSameSite(http.SameSiteLaxMode)
		c.Header("x-access-token", tokens[0])
		c.SetCookie("todolist_refresh", tokens[1], int(time.Hour*24), "/", "", false, true)

		c.Next()
	}
}
