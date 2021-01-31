package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID        uint      `json:"id,primary_key"`
	Email     string    `json:"email" validate:"required" gorm:"uniqueIndex;size:255"`
	Password  string    `json:"password" validate:"required" sql:"password"`
	Username  string    `json:"username" sql:"username"`
	CreatedAt time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt time.Time `json:"updatedat" sql:"updatedat"`
}

func (u User) HashAndSalt(pwd []byte) (string, error) {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        return "", err
    }

    return string(hash), nil
}

func (u User) CompareHashAndPassword(hashedPw, pw []byte) (bool, error) {
	// Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPw)
    err := bcrypt.CompareHashAndPassword(byteHash, pw)
    if err != nil {
        return false, err
    }
    
    return true, nil
}