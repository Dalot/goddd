package factory

import (
	"github.com/Dalot/goddd/internal/app/adapter/repository"
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/brianvoe/gofakeit"
)

// Project is the factory of domain.Project
type Project struct{}

var (
	projectRepository = repository.Project{}
	userRepository    = repository.User{}
)

// Generate generates domain.Project from primitives
func (of Project) Generate() domain.Project {
	factoryUser := User{}
	user := factoryUser.Generate()
	user = userRepository.Save(user)

	project := domain.Project{
		Name:   gofakeit.Name(),
		UserID: user.ID,
	}

	return project
}
