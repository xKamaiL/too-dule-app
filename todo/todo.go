package todo

import (
	"context"
	"errors"
	"fmt"
	"github.com/xkamail/too-dule-app/member"
	"github.com/xkamail/too-dule-app/pkg/config"
	"log"
	"time"
)

type Service struct {
	repo       *Repository
	cfg        *config.Config
	memberRepo *member.Repository
	// logging
	log interface{}
}

func NewTodo() *Service {
	memberRepo := member.NewRepo()
	return &Service{
		repo:       &Repository{},
		cfg:        config.Load(),
		memberRepo: memberRepo,
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
	m, err := t.memberRepo.FindByID(ctx, memberID)
	// member not found
	if err != nil {
		fmt.Println(err)
		return errors.New("member not found")
	}

	return t.repo.UpdateAssignIDByID(ctx, todoID, &m.ID)
}

func (t Service) RemoveAssign(ctx context.Context, todoID string) error {
	return t.repo.UpdateAssignIDByID(ctx, todoID, nil)
}

type UpdateParam struct {
	Title   string `json:"title" validate:"required,min=10,max=255"`
	Content string `json:"content" validate:"required,min=1,max=2000"`
}

func (t Service) Update(ctx context.Context, todoID string, params UpdateParam) (*Todo, error) {
	_, err := t.repo.UpdateByID(ctx, todoID, params.Title, params.Content)
	if err != nil {
		return nil, err
	}
	todo, err := t.repo.FindByID(ctx, todoID)
	return todo, err
}

func (t Service) ChangeStatus(ctx context.Context, todoID string, status bool) error {
	return t.repo.UpdateStatusByID(ctx, todoID, status)
}
