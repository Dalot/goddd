package repository

import "github.com/Dalot/goddd/internal/app/domain"

// ITask is interface of Task repository
type ITask interface {
	Index() []domain.Task
	IndexByProjectID(projectID uint) []domain.Task
	Save(domain.Task) domain.Task
	//Update(domain.Task)
}
