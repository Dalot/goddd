package model

import (
	"time"
)

// User is model
type User struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
