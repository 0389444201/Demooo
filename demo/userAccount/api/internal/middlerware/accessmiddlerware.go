package middlerware

import (
	"context"
	"demo/userAccount/rediscache"
)

func AccessMiddleware(ctx context.Context, key string, value string) {
	rdb := rediscache.NewRedisCache()
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		ctx.Err()
	}
}
