package postgres

import (
	"github.com/jmoiron/sqlx"
	"v001_onelab/internal/model"
)

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
