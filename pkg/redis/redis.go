package redis

import (
	"context"
	"github.com/acoshift/middleware"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type (
	ctxRedisClientKey struct{}
)

func NewClientContext(ctx context.Context, c *redis.Client) context.Context {
	return context.WithValue(ctx, ctxRedisClientKey{}, c)
}

func Middleware(c *redis.Client) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = NewClientContext(ctx, c)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func GetClient(ctx context.Context) *redis.Client {
	return ctx.Value(ctxRedisClientKey{}).(*redis.Client)
}
