package model

import (
	"database/sql"
	"time"
)

type OrderBook struct {
	ID         int          `json:"id" db:"id"`
	BookID     int          `json:"book_id" db:"book_id"`
	UserID     int          `json:"user_id" db:"user_id"`
	OrderDate  time.Time    `json:"order_date" db:"order_date"`
	ReturnDate sql.NullTime `json:"return_date" db:"return_date"`
	User       User         `json:"user" db:"user"`
	Book       Book         `json:"book" db:"book"`
}

type UserOrderBooks struct {
	ID       uint         `json:"id" db:"id"`
	FullName string       `json:"fullname" db:"fullname"`
	Login    string       `json:"login" db:"login"`
	Book     BookWithDate `json:"book" db:"book"`
}

type UserOrderBooksResponse struct {
	ID       uint           `json:"id"`
	Login    string         `json:"login"`
	FullName string         `json:"full_name"`
	Book     []BookWithDate `json:"book"`
}

func (u *UserOrderBooksResponse) AddBook(book BookWithDate) {
	u.Book = append(u.Book, book)
}

type BookWithDate struct {
	ID          uint         `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	Author      string       `json:"author" db:"author"`
	OrderDate   time.Time    `json:"order_date" db:"order_date"`
	ReturnDate  sql.NullTime `json:"return_date" db:"return_date"`
}
