package model

import (
	"time"
)

// Project is model
type Project struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
