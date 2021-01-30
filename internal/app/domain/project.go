package domain

import "time"

type Project struct {
	ID        uint
	Name      string
	UserID    uint
	CreatedAt time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt time.Time `json:"updatedat" sql:"updatedat"`
}
