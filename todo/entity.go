package todo

import (
	"github.com/xkamail/too-dule-app/member"
	"time"
)

type Todo struct {
	ID           string         `json:"id"`
	OwnerID      string         `json:"owner_id"`
	Content      string         `json:"content"`
	IsActive     bool           `json:"is_active"`
	DueDate      *time.Time     `json:"due_date"`
	CreatedAt    *time.Time     `json:"created_at"`
	Owner        *member.Member `json:"owner"`
	AssignID     *string        `json:"assign_id"`
	AssignMember *member.Member `json:"assign_member"`
}
