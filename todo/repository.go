package todo

import (
	"context"
	"github.com/acoshift/pgsql"
	"github.com/acoshift/pgsql/pgctx"
	"github.com/xkamail/too-dule-app/member"
	"log"
	"time"
)

// type BaseRepository interface {
// 	FindAll() ([]Todo, error)
// 	FindByID(id string) (Todo, error)
// }

type Repository struct {
	// ctx context.Context
}

func (r Repository) FindAll(ctx context.Context) ([]*Todo, error) {
	result := make([]*Todo, 0)

	// language=SQL
	err := pgctx.Iter(ctx, func(scan pgsql.Scanner) error {
		var t Todo
		var m member.Member
		t.Owner = &m
		if err := scan(&t.ID, &t.OwnerID, &t.Content, &t.IsActive, &t.DueDate, &t.CreatedAt, &t.Owner.ID, &t.Owner.Username, &t.Owner.Email, &t.Owner.CreatedAt); err != nil {
			log.Println(err)
			return err
		}
		result = append(result, &t)
		return nil
	}, `select todos.id, todos.owner_id, todos.content, todos.is_active,
       todos.due_date,todos.created_at, members.id,members.username, members.email,members.created_at
		from todos left join members on todos.owner_id = members.id
		`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (r Repository) FindByID(ctx context.Context, id string) (Todo, error) {
	panic("implement me")
}

type createTodoModel struct {
	OwnerID  string
	Title    string
	Content  string
	IsActive bool
	DueDate  *time.Time
}

func (r Repository) Insert(ctx context.Context, data createTodoModel) (insertID string, err error) {
	err = pgctx.RunInTx(ctx, func(ctx context.Context) error {
		// language=SQL
		row := pgctx.QueryRow(ctx, `insert into todos ( owner_id, title, content, is_active, due_date) values ($1,$2,$3,$4,$5) returning id`, data.OwnerID, data.Title, data.Content, data.IsActive, data.DueDate.Format(time.RFC3339))
		return row.Scan(&insertID)
	})
	return insertID, err
}
