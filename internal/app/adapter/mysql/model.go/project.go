package model

import (
	"time"

	"gorm.io/gorm"
)

// Project is model
type Project struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"column:user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

