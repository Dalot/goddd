package repository

import (
	"errors"

	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/domain"
	"gorm.io/gorm"
)

// Project is the repository of domain.Project
type User struct{}

// Index fetches all projects
func (u User) Index() []domain.User {
	db := mysql.Connection()
	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		panic(err)

	}

	return users
}

// Get gets parameter
func (u User) GetByID(ID uint) (domain.User, error) {
	db := mysql.Connection()
	user := domain.User{}
	if err := db.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, err
		} else {
			panic(err)
		}
	}

	return user, nil
}

// Save saves User
func (u User) Save(user domain.User) domain.User {
	db := mysql.Connection()

	if err := db.Save(&user).Error; err != nil {
		panic(err)
	}

	return user
}
