package repository

import "github.com/Dalot/goddd/internal/app/domain"

// ITask is interface of Task repository
type ITask interface {
	Index() []domain.Task
	IndexByProjectID(projectID uint) []domain.Task
	GetByID(id uint) (domain.Task, error)
	Save(domain.Task) domain.Task
	Finish(domain.Task) domain.Task
	//Update(domain.Task)
}
