package repository

import "github.com/Dalot/goddd/internal/app/domain"

// IProject is interface of project repository
type IProject interface {
	Index() []domain.Project
	Save(domain.Project) domain.Project
	//Update(domain.Project)
}
