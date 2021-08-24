package todo

import (
	"context"
	"github.com/xkamail/too-dule-app/pkg/config"
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
	Title   string `json:"title" validate:"required,min=10,max=255"`
	Content string `json:"content" validate:"required,min=1,max=2000"`
	OwnerID string `json:"owner_id"`
}

func (t Service) Create(ctx context.Context, param CreateParam) (todoID string, err error) {

	return "", nil
}

// nothing
type ListParam struct {
}

func (t Service) List(ctx context.Context, param ListParam) ([]*Todo, error) {
	// query, pagination , search
	return t.repo.FindAll(ctx)
}
