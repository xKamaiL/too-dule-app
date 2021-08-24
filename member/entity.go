package member

import (
	"time"
)

type Member struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
}
