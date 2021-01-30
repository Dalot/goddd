package usecase

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

func Projects(projectRepository repository.IProject) []domain.Project {
	return projectRepository.Index()
}
