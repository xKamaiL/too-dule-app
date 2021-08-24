package member

import "context"

type (
	ctxMemberKey struct{}
)

func NewMemberContext(ctx context.Context, member Member) context.Context {
	return context.WithValue(ctx, ctxMemberKey{}, member)
}

func GetMemberFromContext(ctx context.Context) (Member, error) {
	member := ctx.Value(ctxMemberKey{}).(Member)
	// need to check type
	return member, nil
}
