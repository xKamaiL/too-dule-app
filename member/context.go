package member

import (
	"context"
	"errors"
	"fmt"
)

type (
	ctxMemberKey struct{}
)

func NewMemberContext(ctx context.Context, member Member) context.Context {
	return context.WithValue(ctx, ctxMemberKey{}, member)
}

func GetMemberFromContext(ctx context.Context) (*Member, error) {
	fmt.Println(ctx.Value(ctxMemberKey{}))
	member, ok := ctx.Value(ctxMemberKey{}).(Member)
	if !ok {
		return nil, errors.New("failed to get member")
	}
	return &member, nil
}
