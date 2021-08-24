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

// no struct
func (r Repository) Insert(ctx context.Context, d struct {
	Username string
	Password string
	Email    string
}) (userId string, err error) {
	// language=SQL
	row := pgctx.QueryRow(ctx, `insert into members (username, password, email) VALUES ($1,$2,$3)`, d.Username, d.Password, d.Email)
	err = row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}
