package service

import (
	"v001_onelab/configs"
	"v001_onelab/internal/repository"
)

type Service struct {
	User IUser
	Book IBook
}

func New(repo *repository.Repository, config *configs.Config) *Service {
	user := NewUser(repo.User, config.HMACSecret)
	book := NewBook(repo.Book)
	return &Service{
		User: user,
		Book: book,
	}
}
