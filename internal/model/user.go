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

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Login    string `json:"login"`
}

type UpdateUser struct {
	FullName string `json:"fullName"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Login           string `json:"login"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

var ErrorNotFound = errors.New("user not found")
var ErrorAlreadyExist = errors.New("user with this login already exist")
