package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"v001_onelab/internal/model"
)

type IUserRepository interface {
	Create(user model.UserInput) error
	GetByID(id int) (model.UserResponse, error)
	GetByLogin(login string) (model.User, error)
	GetAll() ([]model.UserResponse, error)
	Delete(id int) error
	Update(user model.UserResponse) error
	ChangePassword(user model.ChangePassword) error
}

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Create(user model.UserInput) error {
	query, err := u.db.Preparex("INSERT INTO users(fullName,login, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer query.Close()

	if _, err := query.Exec(user.FullName, user.Login, user.Password); err != nil {
		if err.(*pq.Error).Constraint == "users_login_key" {
			return model.ErrorAlreadyExist
		}
		return err
	}
	return nil
}

func (u User) GetByID(id int) (model.UserResponse, error) {
	query, err := u.db.Preparex("SELECT id,fullName,login FROM users WHERE id = $1")
	if err != nil {
		return model.UserResponse{}, err
	}
	defer query.Close()

	var user model.UserResponse
	err = query.Get(&user, id)
	//check user is found
	if err == sql.ErrNoRows {
		return model.UserResponse{}, model.ErrorNotFound
	} else if err != nil {
		return model.UserResponse{}, err
	}
	return user, nil
}

func (u User) GetByLogin(login string) (model.User, error) {
	query, err := u.db.Preparex("SELECT id,fullName,login,password FROM users WHERE login = $1")
	if err != nil {
		return model.User{}, err
	}
	defer query.Close()

	var user model.User
	err = query.Get(&user, login)
	//check user is found
	if err == sql.ErrNoRows {
		return model.User{}, model.ErrorNotFound
	} else if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u User) GetAll() ([]model.UserResponse, error) {
	var users []model.UserResponse
	if err := u.db.Select(&users, "SELECT id,fullName,login FROM users WHERE isdeleted = false"); err != nil {
		return []model.UserResponse{}, err
	}
	return users, nil
}

func (u *User) Delete(id int) error {
	query, err := u.db.Preparex("UPDATE users SET isdeleted = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	if _, err := query.Exec(true, id); err != nil {
		return err
	}
	return nil
}

func (u User) Update(user model.UserResponse) error {
	_, err := u.db.Exec("UPDATE users SET fullname = $1, login = $2 WHERE id = $3", user.FullName, user.Login, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u User) ChangePassword(user model.ChangePassword) error {
	_, err := u.db.Exec("UPDATE users SET password = $1 WHERE login = $2", user.NewPassword, user.Login)
	if err != nil {
		return err
	}
	return nil
}
