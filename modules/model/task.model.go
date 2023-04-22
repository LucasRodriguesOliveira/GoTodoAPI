package model

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
