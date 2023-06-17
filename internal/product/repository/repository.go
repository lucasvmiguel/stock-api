// repository package is responsible for the data access layer of the product domain
package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// product repository
type Repository struct {
	dbClient *gorm.DB
}

var (
	// error when db client is nil
	ErrNilDBClient = errors.New("DB client cannot be nil")
)

// creates a new product repository
func NewRepository(dbClient *gorm.DB) (*Repository, error) {
	if dbClient == nil {
		return nil, ErrNilDBClient
	}

	return &Repository{dbClient}, nil
}
