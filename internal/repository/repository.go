package repository

import "v001_onelab/pkg/database/postgres"

type Repository struct {
	User IUserRepository
}

func New(db *postgres.DatabasePSQL) *Repository {
	user := NewUser(db.DB)
	return &Repository{
		User: user,
	}
}
