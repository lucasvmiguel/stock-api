package starter

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/pkg/env"
)

// config is a struct that holds all config related info
type config struct {
	ENV                    env.Environment
	DBPort                 string
	DBHost                 string
	DBName                 string
	DBUser                 string
	DBPassword             string
	Port                   string
	ServiceName            string
	PaginationDefaultLimit int
}

// loads configuration from the environment variables
func loadConfig() (config, error) {
	paginationDefaultLimitStr := os.Getenv("PAGINATION_DEFAULT_LIMIT")
	paginationDefaultLimit, err := strconv.Atoi(paginationDefaultLimitStr)
	if err != nil {
		return config{}, errors.Wrap(err, "failed to convert PAGINATION_DEFAULT_LIMIT env var to int")
	}

	return config{
		ENV:                    env.Environment(os.Getenv("ENV")),
		Port:                   os.Getenv("PORT"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBName:                 os.Getenv("DB_NAME"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBPort:                 os.Getenv("DB_PORT"),
		ServiceName:            os.Getenv("SERVICE_NAME"),
		PaginationDefaultLimit: paginationDefaultLimit,
	}, nil
}
