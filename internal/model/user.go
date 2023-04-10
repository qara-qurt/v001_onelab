package model

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type User struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullname"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserInput struct {
	FullName string `json:"fullName" validate:"required,gte=2"`
	Login    string `json:"login" validate:"required,gte=4"`
	Password string `json:"password" validate:"required,gte=6"`
}

func (u UserInput) Validate() error {
	return validate.Struct(u)
}

type SignInInput struct {
	Login    string `json:"login" validate:"required,gte=4"`
	Password string `json:"password" validate:"required,gte=6"`
}

func (s SignInInput) Validate() error {
	return validate.Struct(s)
}

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName" validate:"required,gte=2"`
	Login    string `json:"login" validate:"required,gte=4"`
}

func (u UserResponse) Validate() error {
	return validate.Struct(u)
}

type UpdateUser struct {
	FullName string `json:"fullName" `
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Login           string `json:"login" validate:"required,gte=4"`
	CurrentPassword string `json:"currentPassword" validate:"required,gte=6"`
	NewPassword     string `json:"newPassword" validate:"required,gte=6"`
}

func (c ChangePassword) Validate() error {
	return validate.Struct(c)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Token struct {
	Token string `json:"token"`
}

func NewErrorResponse(err string) ErrorResponse {
	return ErrorResponse{Error: err}
}

var ErrorNotFound = errors.New("user not found")
var ErrorAlreadyExist = errors.New("user with this login already exist")
var ErrorPassword = errors.New("password is not correct")
