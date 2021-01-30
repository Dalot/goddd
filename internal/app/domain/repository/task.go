package repository

import "github.com/Dalot/goddd/internal/app/domain"

// ITask is interface of Task repository
type ITask interface {
	Index() []domain.Task
	IndexByProjectID(projectID uint) []domain.Task
	GetByID(id uint) (domain.Task, error)
	Create(domain.Task) domain.Task
	Save(domain.Task) domain.Task
	Delete(id uint)
}
