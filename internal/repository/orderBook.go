package repository

import (
	"github.com/jmoiron/sqlx"
	"v001_onelab/internal/model"
)

type IOrderBookRepository interface {
	GetOrderBooks() ([]model.OrderBook, error)
	GetOrderUserBooks(isLastMounth bool) ([]model.UserOrderBooks, error)
}

type OrderBook struct {
	db *sqlx.DB
}

func NewOrderBook(db *sqlx.DB) *OrderBook {
	return &OrderBook{
		db: db,
	}
}

func (o OrderBook) GetOrderBooks() ([]model.OrderBook, error) {
	var books []model.OrderBook
	query := `
		SELECT
			obh.id,
			obh.user_id,
			u.fullname AS "user.fullname",
			u.login AS "user.login",
			b.id AS "book.id",
			b.name AS "book.name",
			b.description AS "book.description",
			b.author AS "book.author"
		FROM
			order_book_history obh
			JOIN users u ON u.id = obh.user_id
			JOIN books b ON b.id = obh.book_id
	`
	if err := o.db.Select(&books, query); err != nil {
		return []model.OrderBook{}, err
	}
	return books, nil
}

func (o OrderBook) GetOrderUserBooks(isLastMounth bool) ([]model.UserOrderBooks, error) {
	var userBooks []model.UserOrderBooks
	query := `
		SELECT
			obh.user_id AS id,
			u.fullname,
			u.login,
			b.id AS "book.id",
			b.name AS "book.name",
			b.description AS "book.description",
			b.author AS "book.author",
			obh.order_date AS "book.order_date",
			obh.return_date AS "book.return_date"
		FROM
			order_book_history obh
			JOIN users u ON u.id = obh.user_id
			JOIN books b ON b.id = obh.book_id
	`
	if isLastMounth {
		query += `WHERE obh.order_date >= DATE_TRUNC('month', NOW())`
	}
	if err := o.db.Select(&userBooks, query); err != nil {
		return []model.UserOrderBooks{}, err
	}
	return userBooks, nil
}
