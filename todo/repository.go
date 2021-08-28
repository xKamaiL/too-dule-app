package todo

import (
	"context"
	"database/sql"
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

//
func NewRepo() *Repository {
	return &Repository{}
}

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

func (r Repository) FindByID(ctx context.Context, id string) (*Todo, error) {
	result := new(Todo)
	// language=SQL
	row := pgctx.QueryRow(ctx, `select todos.id, todos.owner_id, todos.content, todos.is_active,
       todos.due_date,todos.created_at
		from todos where todos.id = $1 limit 1`, id)

	if err := row.Scan(&result.ID, &result.OwnerID, &result.Content, &result.IsActive, &result.DueDate, &result.CreatedAt); err != nil {
		return nil, err
	}

	return result, nil
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

// no struct
func (r Repository) UpdateByID(ctx context.Context, id string, title string, content string) (sql.Result, error) {
	// language=SQL
	result, err := pgctx.Exec(ctx, "update todos set content = $1,title = $2 where id = $3", content, title, id)
	return result, err
}

func (r Repository) UpdateStatusByID(ctx context.Context, id string, status bool) error {
	// language=SQL
	_, err := pgctx.Exec(ctx, `update todos set is_active = $1 where id = $2`, status, id)
	return err
}

func (r Repository) UpdateAssignIDByID(ctx context.Context, id string, assignID *string) error {
	// language=SQL
	_, err := pgctx.Exec(ctx, `update todos set assign_id = $1 where id = $2`, assignID, id)
	return err
}

func (r Repository) DeleteByID(ctx context.Context, id string) error {
	// language=SQL
	_, err := pgctx.Exec(ctx, `delete from todos where id = $1`, id)
	return err
}
