package service

import (
	"context"
	"v001_onelab/configs"
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type IOrderBook interface {
	GetOrderBooks() ([]model.OrderBook, error)
	GetOrderUserBooks() ([]model.UserOrderBooksResponse, error)
	GetOrderUserBooksLastMounth() ([]model.UserOrderBooksResponse, error)
}

type IBook interface {
	Create(book model.BookInput) error
	GetAll() ([]model.Book, error)
}

type IUser interface {
	Create(user model.UserInput) (int, error)
	SignIn(user model.SignInInput) (string, error)
	ParseToken(ctx context.Context, token string) (uint, error)
	GetByID(id int) (model.UserResponse, error)
	GetByLogin(login string) (model.User, error)
	GetAll() ([]model.UserResponse, error)
	Delete(id int) error
	Update(user model.UserResponse) error
	ChangePassword(user model.ChangePassword) error
}

type Service struct {
	User      IUser
	Book      IBook
	OrderBook IOrderBook
}

func New(repo *repository.Repository, config *configs.Config) *Service {
	user := NewUser(repo.User, config.HMACSecret)
	book := NewBook(repo.Book)
	orderBook := NewOrderBook(repo.OrderBook)
	return &Service{
		User:      user,
		Book:      book,
		OrderBook: orderBook,
	}
}
