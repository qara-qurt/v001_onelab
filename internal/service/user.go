package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository"
)

type IUser interface {
	Create(user model.User) error
	GetByID(id int) (model.UserResponse, error)
	GetByLogin(login string) (model.User, error)
	GetAll() ([]model.UserResponse, error)
	Delete(id int) error
	Update(user model.UserResponse) error
	ChangePassword(user model.ChangePassword) error
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
	hashPass, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass
	return u.repo.Create(user)
}

func (u User) GetByID(id int) (model.UserResponse, error) {
	return u.repo.GetByID(id)
}

func (u User) GetByLogin(login string) (model.User, error) {
	return u.repo.GetByLogin(login)
}

func (u User) GetAll() ([]model.UserResponse, error) {
	return u.repo.GetAll()
}

func (u User) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u User) Update(user model.UserResponse) error {
	return u.repo.Update(user)
}

func (u User) ChangePassword(user model.ChangePassword) error {
	res, err := u.repo.GetByLogin(user.Login)
	if err != nil {
		return err
	}
	ok := checkPasswordHash(user.CurrentPassword, res.Password)
	if !ok {
		return errors.New("password is not correct")
	}
	hashPass, err := hashPassword(user.NewPassword)
	if err != nil {
		return err
	}
	user.NewPassword = hashPass
	return u.repo.ChangePassword(user)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
