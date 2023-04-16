package repository

import (
	"v001_onelab/configs"
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository/postgres"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type IOrderBookRepository interface {
	GetOrderBooks() ([]model.OrderBook, error)
	GetOrderUserBooks(isLastMounth bool) ([]model.UserOrderBooks, error)
}

type IBookRepository interface {
	Create(book model.BookInput) error
	GetAll() ([]model.Book, error)
}

type IUserRepository interface {
	Create(user model.UserInput) (int, error)
	GetByID(id int) (model.UserResponse, error)
	GetByLogin(login string) (model.User, error)
	GetAll() ([]model.UserResponse, error)
	Delete(id int) error
	Update(user model.UserResponse) error
	ChangePassword(user model.ChangePassword) error
}

type Repository struct {
	User      IUserRepository
	Book      IBookRepository
	OrderBook IOrderBookRepository
}

func New(config *configs.Config) (*Repository, error) {
	db, err := postgres.NewDatabasePSQL(config)
	if err != nil {
		return nil, err
	}

	user := postgres.NewUser(db.DB)
	book := postgres.NewBook(db.DB)
	orderBook := postgres.NewOrderBook(db.DB)
	return &Repository{
		User:      user,
		Book:      book,
		OrderBook: orderBook,
	}, nil
}
