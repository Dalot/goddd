package model

import (
	"time"
)

// Project is model
type Task struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	ProjectID   string `gorm:"column:project_id"`
	Description string
	Status      string
	CreatedAt   time.Time
	FinishedAt  string
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
