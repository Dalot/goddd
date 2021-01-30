package domain

import "time"

type Task struct {
	ID          uint
	Name        string
	Description string
	ProjectID   uint
	Status      string
	FinishedAt  string `json:"finished_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_dat"`
}