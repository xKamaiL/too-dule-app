package member

import (
	"context"
	"errors"
)

type (
	ctxMemberKey struct{}
)

func NewMemberContext(ctx context.Context, member Member) context.Context {
	return context.WithValue(ctx, ctxMemberKey{}, member)
}

func GetMemberFromContext(ctx context.Context) (*Member, error) {
	member, ok := ctx.Value(ctxMemberKey{}).(Member)
	if !ok {
		return nil, errors.New("failed to get member")
	}
	return &member, nil
}
