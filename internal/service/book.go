package service

import (
	"v001_onelab/internal/model"
	"v001_onelab/internal/repository"
)

type Book struct {
	repo repository.IBookRepository
}

func NewBook(repo repository.IBookRepository) *Book {
	return &Book{
		repo: repo,
	}
}

func (b Book) Create(book model.BookInput) error {
	return b.repo.Create(book)
}

func (b Book) GetAll() ([]model.Book, error) {
	return b.repo.GetAll()
}
