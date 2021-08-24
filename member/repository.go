package member

import (
	"context"
	"github.com/acoshift/pgsql"
	"github.com/acoshift/pgsql/pgctx"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"log"
	"time"
)

type Repository struct {
}

func (r Repository) FindByID(ctx context.Context, memberID string) (*Member, error) {
	var member Member
	// language=SQL
	row := pgctx.QueryRow(ctx, `select id, username, email, created_at,password from members where id = $1`, memberID)
	err := row.Scan(&member.ID, &member.Username, &member.Email, &member.CreatedAt, &member.Password)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r Repository) FindByUsername(ctx context.Context, username string) (*Member, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("FindByUsername took %s", elapsed)
	}()

	var member Member
	// language=SQL
	row := pgctx.QueryRow(ctx, `select id, username, email, created_at,password from members where username = $1`, username)
	err := row.Scan(&member.ID, &member.Username, &member.Email, &member.CreatedAt, &member.Password)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r Repository) FindByEmail(ctx context.Context, email string) (*Member, error) {
	var member Member
	// language=SQL
	row := pgctx.QueryRow(ctx, `select id, username, email, created_at,password from members where email = $1`, email)
	err := row.Scan(&member.ID, &member.Username, &member.Email, &member.CreatedAt, &member.Password)
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
	password, err := utils.HashPassword(d.Password)
	if err != nil {
		return "", err
	}
	// language=SQL
	row := pgctx.QueryRow(ctx, `insert into members (username, password, email) VALUES ($1,$2,$3) returning id`, d.Username, password, d.Email)
	err = row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (r Repository) FindAll(ctx context.Context) ([]*Member, error) {
	list := make([]*Member, 0)
	// language=SQL
	err := pgctx.Iter(ctx, func(scan pgsql.Scanner) error {
		var m Member
		if err := scan(&m.ID, &m.Username, &m.Email); err != nil {
			return err
		}
		list = append(list, &m)
		return nil
	}, `select id,username,email from members order by created_at desc `)
	return list, err
}
