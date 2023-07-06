// service package is responsible for the business logic of the product domain
package service

import (
	"context"
	"errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// product service that manages different features for a product
type Service struct {
	repository Repository
	transactor Transactor
}

// Transactor is the interface that wraps the required transactor methods
type Transactor interface {
	Begin(ctx context.Context) context.Context
	Rollback(ctx context.Context)
	Commit(ctx context.Context)
}

// repository interface that can be implemented by any kind of storage
type Repository interface {
	Create(ctx context.Context, product entity.Product) (*entity.Product, error)
	GetAll(ctx context.Context) ([]*entity.Product, error)
	GetByID(ctx context.Context, id int) (*entity.Product, error)
	GetPaginated(ctx context.Context, cursor int, limit int) ([]*entity.Product, error)
	UpdateByID(ctx context.Context, id int, product entity.Product) (*entity.Product, error)
	DeleteByID(ctx context.Context, id int) (*entity.Product, error)
}

// NewServiceArgs is the arguments struct for NewService function
type NewServiceArgs struct {
	Repository Repository
	Transactor Transactor
}

// creates a new product service
func NewService(args NewServiceArgs) (*Service, error) {
	if args.Repository == nil {
		return nil, errors.New("repository is required")
	}

	if args.Transactor == nil {
		return nil, errors.New("transactor is required")
	}

	return &Service{
		repository: args.Repository,
		transactor: args.Transactor,
	}, nil
}
