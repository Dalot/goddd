package model

import (
	"time"

	"gorm.io/gorm"
)

// User is model
type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string
	Password     string
	Username     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

