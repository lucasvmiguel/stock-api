// repository package is responsible for the data access layer of the product domain
package repository

import (
	"context"

	"github.com/lucasvmiguel/stock-api/pkg/transactor"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	idColumnName = "id"
)

// product repository
type Repository struct {
	db *gorm.DB
}

var (
	// error when db client is nil
	ErrNilDBClient = errors.New("DB client cannot be nil")
)

// creates a new product repository
func NewRepository(db *gorm.DB) (*Repository, error) {
	if db == nil {
		return nil, ErrNilDBClient
	}

	return &Repository{db}, nil
}

func (r *Repository) run(ctx context.Context) *gorm.DB {
	transaction := transactor.DBTransaction(ctx)

	if transaction != nil {
		return transaction
	}

	return r.db
}
