package member

import (
	"context"
	"github.com/acoshift/pgsql/pgctx"
)

type Repository struct {
}

func (r Repository) FindByID(ctx context.Context, memberID string) (*Member, error) {
	var member Member
	// language=SQL
	row := pgctx.QueryRow(ctx, `select id, username, email, created_at from members where id = $1`, memberID)
	err := row.Scan(&member.ID, &member.Username, &member.Email, &member.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &member, nil
}
