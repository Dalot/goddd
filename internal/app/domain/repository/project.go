package repository

import "github.com/Dalot/goddd/internal/app/domain"

// IProject is interface of project repository
type IProject interface {
	Index() []domain.Project
	IndexByUserID(userId uint) []domain.Project
	Save(domain.Project) domain.Project
	Delete(id uint)
	GetByID(id uint) (domain.Project, error)
	//Update(domain.Project)
}
