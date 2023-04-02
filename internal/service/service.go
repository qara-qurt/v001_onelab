package service

import (
	"v001_onelab/internal/repository"
)

type Service struct {
	User repository.IUserRepository
}

func New(repo *repository.Repository) *Service {
	user := NewUser(repo.User)
	return &Service{
		User: user,
	}
}
