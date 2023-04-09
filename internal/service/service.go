package service

import (
	"v001_onelab/configs"
	"v001_onelab/internal/repository"
)

type Service struct {
	User IUser
}

func New(repo *repository.Repository, config *configs.Config) *Service {
	user := NewUser(repo.User, config.HMACSecret)
	return &Service{
		User: user,
	}
}
