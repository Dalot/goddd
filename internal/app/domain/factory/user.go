package factory

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/brianvoe/gofakeit"
)

// User is the factory of domain.Project
type User struct{}

// Generate generates domain.Project from primitives
func (of User) Generate() domain.User {
	pw := []byte("123123123")

	user := domain.User{
		Email:    gofakeit.Email(),
		Username: gofakeit.Name(),
	}

	hash, err := user.HashAndSalt(pw)
	if err != nil {
		panic(err)
	}
	user.Password = hash
	
	return user
}
