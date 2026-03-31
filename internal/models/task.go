package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,min=3"`
	Description string    `json:"description" validate:"required"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
