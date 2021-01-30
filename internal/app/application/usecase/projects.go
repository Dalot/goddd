package usecase

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

type CreateProjectArgs struct {
	ProjectID         string
	Name              string
	UserID            uint
	ProjectRepository repository.IProject
	UserRepository    repository.IUser
}

func CreateProject(args CreateProjectArgs) (domain.Project, error) {
	user, err := args.UserRepository.GetByID(args.UserID)

	if err != nil {
		return domain.Project{}, err
	}

	project := domain.Project{
		Name:   args.Name,
		UserID: user.ID,
	}

	args.ProjectRepository.Save(project)
	return project, nil
}

func Projects(projectRepository repository.IProject, userID uint) []domain.Project {
	if userID == 0 {
		return projectRepository.Index()
	} else {
		return projectRepository.IndexByUserID(userID)
	}
}

func GetProjectByID(projectRepository repository.IProject, userID uint) (domain.Project, error) {
	proj, err := projectRepository.GetByID(userID)
	if err != nil {
		return proj, err
	}
	return proj, nil
}

func DeleteProject(projectRepository repository.IProject, projectID uint) error {
	// TODO: if project not found, return an error
	_, err := projectRepository.GetByID(projectID)
	if err != nil {
		return err
	}
	projectRepository.Delete(projectID)

	return nil
}
