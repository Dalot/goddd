package usecase

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

type CreateProjectArgs struct {
	ProjectID         uint
	Name              string
	UserID            uint
	ProjectRepository repository.IProject
	UserRepository    repository.IUser
}

type UpdateProjectArgs struct {
	ProjectID         uint
	Name              string
	ProjectRepository repository.IProject
}

func CreateProject(args CreateProjectArgs) (domain.Project, error) {
	user, err := args.UserRepository.GetByID(args.UserID)
	if err != nil {
		return domain.Project{}, err
	}

	proj := domain.Project{
		Name:   args.Name,
		UserID: user.ID,
	}

	proj = args.ProjectRepository.Create(proj)
	return proj, nil
}

func UpdateProject(args UpdateProjectArgs) (domain.Project, error) {
	proj, err := args.ProjectRepository.GetByID(args.ProjectID)
	if err != nil {
		return proj, err
	}

	proj.Name = args.Name

	proj = args.ProjectRepository.Save(proj)
	return proj, nil
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
	_, err := projectRepository.GetByID(projectID)
	if err != nil {
		return err
	}
	projectRepository.Delete(projectID)

	return nil
}
