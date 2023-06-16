package handler

import (
	"errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

var (
	// field id is used as url param for different handlers (eg: get-by-id handler)
	FieldID = "id"
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
	Create(product entity.Product) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	UpdateByID(id uint, product entity.Product) (*entity.Product, error)
	DeleteByID(id uint) (*entity.Product, error)
}

// product handler that has methods to handle different types of http requests
type Handler struct {
	service Service
}

// creates a new product handler
func NewHandler(service Service) (*Handler, error) {
	if service == nil {
		return nil, ErrNilSercice
	}

	return &Handler{service}, nil
}
