package db

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description,"`
	Deadline    time.Time `json:"deadline"`
}

type CreateTodoDto struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type UpdateTodoDto struct {
	CreateTodoDto
}
