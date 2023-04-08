package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"v001_onelab/configs"
)

type DatabasePSQL struct {
	DB *sqlx.DB
}

func NewDatabasePSQL(conf *configs.Config) (*DatabasePSQL, error) {
	db, err := sqlx.Open("postgres", conf.PgURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DatabasePSQL{
		DB: db,
	}, nil
}
