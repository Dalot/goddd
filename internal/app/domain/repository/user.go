package repository

import "github.com/Dalot/goddd/internal/app/domain"

// IProject is interface of project repository
type IUser interface {
	Index() []domain.User
	GetByID(ID string) domain.User
	//Update(domain.User)
}
