package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PORT string
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")
	return &Config{
		PORT: port,
	}, nil
}
