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
	// error when service is nil
	ErrNilSercice = errors.New("service cannot be nil")
	// error when json body is not valid
	ErrInvalidJSONBody = errors.New("invalid JSON body")
	// error when product was not found
	ErrNotFound = errors.New("product not found")
	// error internal server error
	ErrInternalServerError = errors.New("internal server error")
)

// service interface to run different features
type Service interface {
	UpdateByID(id uint, product entity.Product) (*entity.Product, error)
	Create(product entity.Product) (*entity.Product, error)
}

// repository interface that can be implemented by any kind of storage
type Repository interface {
	GetAll() ([]*entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	DeleteByID(id uint) (*entity.Product, error)
}

// product handler that has methods to handle different types of http requests
type Handler struct {
	repository Repository
	service    Service
}

// creates a new product handler
func NewHandler(repository Repository, service Service) (*Handler, error) {
	if repository == nil {
		return nil, ErrNilRepository
	}

	if service == nil {
		return nil, ErrNilSercice
	}

	return &Handler{repository, service}, nil
}
