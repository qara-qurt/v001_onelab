package model

import (
	"errors"
)

type User struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateUser struct {
	FullName string `json:"fullName"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var ErrorNotFound = errors.New("user not found")
var ErrorAlreadyExist = errors.New("user with this login already exist")
