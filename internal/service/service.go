package service

import (
	"v001_onelab/configs"
	"v001_onelab/internal/repository"
)

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
