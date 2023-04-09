package repository

import "v001_onelab/pkg/database/postgres"

type Repository struct {
	User IUserRepository
	Book IBookRepository
}

func New(db *postgres.DatabasePSQL) *Repository {
	user := NewUser(db.DB)
	book := NewBook(db.DB)
	return &Repository{
		User: user,
		Book: book,
	}
}
