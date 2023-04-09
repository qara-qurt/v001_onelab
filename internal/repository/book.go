package repository

import (
	"github.com/jmoiron/sqlx"
	"v001_onelab/internal/model"
)

type IBookRepository interface {
	Create(book model.BookInput) error
	GetAll() ([]model.Book, error)
	GetOrderBooks() ([]model.OrderBook, error)
}

type Book struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) *Book {
	return &Book{
		db: db,
	}
}

func (b Book) Create(book model.BookInput) error {
	query, err := b.db.Preparex("INSERT INTO books(name,description,author) VALUES ($1,$2,$3)")
	if err != nil {
		return err
	}
	if _, err := query.Exec(book.Name, book.Description, book.Author); err != nil {
		return err
	}
	return nil
}

func (b Book) GetAll() ([]model.Book, error) {
	var books []model.Book
	if err := b.db.Select(&books, "SELECT id,name,description,author FROM books"); err != nil {
		return []model.Book{}, err
	}
	return books, nil
}

func (b Book) GetOrderBooks() ([]model.OrderBook, error) {
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
	if err := b.db.Select(&books, query); err != nil {
		return []model.OrderBook{}, err
	}
	return books, nil
}
