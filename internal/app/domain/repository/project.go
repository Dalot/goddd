package repository

import "github.com/Dalot/goddd/internal/app/domain"

// IProject is interface of project repository
type IProject interface {
	Index() []domain.Project
	IndexByUserID(userId uint) []domain.Project
	GetByID(id uint) (domain.Project, error)
	Create(domain.Project) domain.Project
	Save(domain.Project) domain.Project
	Delete(id uint)
	//Update(domain.Project)
}
