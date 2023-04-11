package repository

import "v001_onelab/pkg/database/postgres"

type Repository struct {
	User      IUserRepository
	Book      IBookRepository
	OrderBook IOrderBookRepository
}

func New(db *postgres.DatabasePSQL) *Repository {
	user := NewUser(db.DB)
	book := NewBook(db.DB)
	orderBook := NewOrderBook(db.DB)
	return &Repository{
		User:      user,
		Book:      book,
		OrderBook: orderBook,
	}
}
