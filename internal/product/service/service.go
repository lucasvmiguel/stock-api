package service

import (
	"errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

var (
	// error when repository is nil
	ErrNilRepository = errors.New("repository cannot be nil")
)

// product service that manages different features for a product
type Service struct {
	repository Repository
}

// repository interface that can be implemented by any kind of storage
type Repository interface {
	Create(product entity.Product) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	GetPaginated(cursor uint, limit uint) ([]*entity.Product, error)
	UpdateByID(id uint, product entity.Product) (*entity.Product, error)
	DeleteByID(id uint) (*entity.Product, error)
}

// creates a new product service
func NewService(repository Repository) (*Service, error) {
	if repository == nil {
		return nil, ErrNilRepository
	}

	return &Service{repository}, nil
}
