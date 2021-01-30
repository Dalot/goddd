package repository

import (
	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/domain"
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
func (u User) GetByID(ID string) domain.User {
	db := mysql.Connection()
	user := domain.User{}
	result := db.First(&user, ID)
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}

// Save saves User
func (u User) Save(user domain.User) domain.User {
	db := mysql.Connection()

	if err := db.Save(user).Error; err != nil {
		panic(err)

	}

	return user
}
