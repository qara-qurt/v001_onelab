package service

import (
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository"
)

type IUser interface {
	Create(user model.User) error
	GetByID(id int) (model.User, error)
	GetAll() ([]model.User, error)
	Delete(id int) error
	Update(user model.User) (model.User, error)
}

type User struct {
	repo repository.IUserRepository
}

func NewUser(repo repository.IUserRepository) *User {
	return &User{
		repo: repo,
	}
}

func (u User) Create(user model.User) error {
	users, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Login == user.Login {
			return model.ErrorAlreadyExist
		}
	}
	return u.repo.Create(user)
}

func (u User) GetByID(id int) (model.User, error) {
	return u.repo.GetByID(id)
}

func (u User) GetAll() ([]model.User, error) {
	return u.repo.GetAll()
}

func (u User) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u User) Update(user model.User) (model.User, error) {
	return u.repo.Update(user)
}
