package todo

import (
	"context"
	"github.com/xkamail/too-dule-app/pkg/config"
	"log"
	"time"
)

type Service struct {
	repo *Repository
	cfg  *config.Config
	// logging
	log interface{}
}

func NewTodo() *Service {
	return &Service{
		repo: &Repository{},
		cfg:  config.Load(),
	}
}

type CreateParam struct {
	Title   string     `json:"title" validate:"required,min=10,max=255"`
	Content string     `json:"content" validate:"required,min=1,max=2000"`
	OwnerID string     `json:"owner_id"`
	DueDate *time.Time `json:"due_date" validate:"required"`
}

func (t Service) Create(ctx context.Context, param CreateParam) (todoID string, err error) {
	id, err := t.repo.Insert(ctx, createTodoModel{
		OwnerID:  param.OwnerID,
		Content:  param.Content,
		Title:    param.Title,
		IsActive: false,
		DueDate:  param.DueDate,
	})
	log.Println(err)
	return id, err
}

// nothing
type ListParam struct {
}

func (t Service) List(ctx context.Context, param ListParam) ([]*Todo, error) {
	// query, pagination , search
	return t.repo.FindAll(ctx)
}

func (t Service) AssignToMember(ctx context.Context, todoID string, memberID string) error {

	return nil
}

func (t Service) RemoveAssign(ctx context.Context, todoID string) error {
	return nil
}

type UpdateParam struct {
	Title   string
	Content string
}

func (t Service) Update(ctx context.Context, params UpdateParam) error {
	return nil
}

func (t Service) ChangeStatus(ctx context.Context, status bool) error {
	return nil
}
