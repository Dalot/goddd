package usecase

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

func Tasks(taskRepository repository.ITask, projectID uint) []domain.Task {
	if projectID == 0 {
		return taskRepository.Index()
	} else {
		return taskRepository.IndexByProjectID(projectID)
	}
}
