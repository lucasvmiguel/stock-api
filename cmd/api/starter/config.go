package starter

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/pkg/env"
)

type config struct {
	DBPort                 string
	DBHost                 string
	DBName                 string
	DBUser                 string
	DBPassword             string
	Port                   string
	ENV                    env.Environment
	PaginationDefaultLimit int
}

func loadConfig() (config, error) {
	paginationDefaultLimitStr := os.Getenv("PAGINATION_DEFAULT_LIMIT")
	paginationDefaultLimit, err := strconv.Atoi(paginationDefaultLimitStr)
	if err != nil {
		return config{}, errors.Wrap(err, "failed to convert PAGINATION_DEFAULT_LIMIT env var to int")
	}

	return config{
		Port:                   os.Getenv("PORT"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBName:                 os.Getenv("DB_NAME"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBPort:                 os.Getenv("DB_PORT"),
		ENV:                    env.Environment(os.Getenv("ENV")),
		PaginationDefaultLimit: paginationDefaultLimit,
	}, nil
}
