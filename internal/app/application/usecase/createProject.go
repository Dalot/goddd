package usecase

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

type CreateProjectArgs struct {
	ProjectID         string
	Name              string
	CreatedBy         string
	ProjectRepository repository.IProject
	UserRepository    repository.IUser
}

func CreateProject(args CreateProjectArgs) domain.Project {
	user := args.UserRepository.GetByID(args.CreatedBy)

	project := domain.Project{
		Name:   args.Name,
		UserID: user.ID,
	}

	args.ProjectRepository.Save(project)
	return project
}
