package ctxutil

import (
	"context"

	"github.com/crazyfrankie/ddd-todolist/backend/pkg/ctxcache"
	"github.com/crazyfrankie/ddd-todolist/backend/types/consts"
)

func MustGetUIDFromCtx(ctx context.Context) int64 {
	data, ok := ctxcache.Get[int64](ctx, consts.SessionDataKeyInCtx)
	if !ok {
		panic("mustGetUIDFromCtx: sessionData is nil")
	}

	return data
}
