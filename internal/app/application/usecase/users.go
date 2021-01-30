package usecase

import (
	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
)

func Users(userRepository repository.IUser) []domain.User {
	return userRepository.Index()
}
