package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"v001_onelab/internal/model"
)
// user нужно хранить в папке postgres | redis | etc 
type IUserRepository interface {
	Create(user model.User) error
	GetByID(id int) (model.User, error)
	GetAll() ([]model.User, error)
	Delete(id int) error
	Update(user model.User) (model.User, error)
}

type User struct {
	db []model.User
}

func NewUser() *User {
	return &User{
		db: []model.User{
			{ID: 1, FullName: "Serikov Dias", Login: "dias002", Password: "qwerty"},
			{ID: 2, FullName: "test2 Dias", Login: "dias003", Password: "qwerty3"},
			{ID: 3, FullName: "test3 Dias", Login: "dias004", Password: "qwerty4"},
			{ID: 4, FullName: "test4 Dias", Login: "dias005", Password: "qwerty5"},
		},
	}
}

func (u *User) Create(user model.User) error {
	if user.FullName == "" {
		return errors.New("user FullName cannot be empty")
	}
	if user.Login == "" {
		return errors.New("user Login cannot be empty")
	}
	if user.Password == "" {
		return errors.New("user Password cannot be empty")
	}
	user.ID = uint(uuid.New().ID())
	u.db = append(u.db, user)
	return nil
}

func (u User) GetByID(id int) (model.User, error) {
	for _, user := range u.db {
		if user.ID == uint(id) {
			return user, nil
		}
	}
	return model.User{}, model.ErrorNotFound
}

func (u User) GetAll() ([]model.User, error) {
	return u.db, nil
}

func (u *User) Delete(id int) error {
	for i, user := range u.db {
		if user.ID == uint(id) {
			u.db = append(u.db[:i], u.db[i+1:]...)
			return nil
		}
	}
	return model.ErrorNotFound
}

func (u User) Update(user model.User) (model.User, error) {
	for i, us := range u.db {
		if us.ID == user.ID {
			u.db[i] = user
			return u.db[i], nil
		}
	}
	return model.User{}, fmt.Errorf("user with ID %d not found", user.ID)
}
