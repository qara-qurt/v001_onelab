package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PORT       string
	Database   string
	PgURL      string
	HMACSecret string
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	//next project I will use "github.com/caarlos0/env/v6" to make code shorter
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	pgURl, ok := os.LookupEnv("PgURL")
	if !ok {
		pgURl = "host=localhost port=5436 user=postgres password=secret dbname=postgres sslmode=disable"
	}
	database, ok := os.LookupEnv("Database")
	if !ok {
		database = "postgres"
	}
	hmacSecret, ok := os.LookupEnv("HMACSecret")
	if !ok {
		hmacSecret = ""
	}
	return &Config{
		PORT:       port,
		PgURL:      pgURl,
		Database:   database,
		HMACSecret: hmacSecret,
	}, nil
}
