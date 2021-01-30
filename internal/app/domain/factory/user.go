package factory

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/brianvoe/gofakeit"
)

// User is the factory of domain.Project
type User struct{}

// Generate generates domain.Project from primitives
func (of User) Generate() domain.User {
	user := domain.User{
		Email:    gofakeit.Email(),
		Username: gofakeit.Name(),
	}

	return user
}
