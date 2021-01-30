package repository

import "github.com/Dalot/goddd/internal/app/domain"

// IProject is interface of project repository
type IUser interface {
	Index() []domain.User
	GetByID(ID uint) (domain.User, error)
	//Update(domain.User)
}
