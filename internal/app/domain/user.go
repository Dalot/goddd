package domain

import "time"

type User struct {
	ID        string    `json:"id,primary_key"`
	Email     string    `json:"email" validate:"required" gorm:"unique"`
	Password  string    `json:"password" validate:"required" sql:"password"`
	Username  string    `json:"username" sql:"username"`
	CreatedAt time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt time.Time `json:"updatedat" sql:"updatedat"`
}