package handler

import (
	"errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

var (
	// field id is used as url param for different handlers (eg: get-by-id handler)
	FieldID = "id"
	// error when repository is nil
	ErrNilRepository = errors.New("repository cannot be nil")
	// error when json body is not valid
	ErrInvalidJSONBody = errors.New("invalid JSON body")
	// error when product was not found
	ErrNotFound = errors.New("product not found")
	// error internal server error
	ErrInternalServerError = errors.New("internal server error")
)

// repository interface that can be implemented by any kind of storage
type Repository interface {
	Create(product entity.Product) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	DeleteByID(id uint) (*entity.Product, error)
	UpdateByID(id uint, product entity.Product) (*entity.Product, error)
}

// product handler that has methods to handle different types of http requests
type Handler struct {
	repository Repository
}

// creates a new product handler
func NewHandler(repository Repository) (*Handler, error) {
	if repository == nil {
		return nil, ErrNilRepository
	}

	return &Handler{repository}, nil
}
