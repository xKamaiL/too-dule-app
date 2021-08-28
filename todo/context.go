package todo

import "context"

type todoCtxKey struct{}

// may be can use on a middleware

func GetFromContext(ctx context.Context) Todo {
	return ctx.Value(todoCtxKey{}).(Todo)
}

func NewFromContext(parent context.Context, v Todo) context.Context {
	return context.WithValue(parent, todoCtxKey{}, v)
}
