package handler

import (
	"errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

var (
	ErrNilRepository   = errors.New("repository cannot be nil")
	ErrInvalidJSONBody = errors.New("invalid JSON body")
	ErrNotFound        = errors.New("product not found")
)

type Repository interface {
	Create(product entity.Product) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	DeleteByID(id uint) (*entity.Product, error)
	UpdateByID(id uint, product entity.Product) (*entity.Product, error)
}

type Handler struct {
	repository Repository
}

func NewHandler(repository Repository) (*Handler, error) {
	if repository == nil {
		return nil, ErrNilRepository
	}

	return &Handler{repository}, nil
}
