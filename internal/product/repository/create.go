package repository

import (
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// creates a product in the database
func (r *Repository) Create(product entity.Product) (*entity.Product, error) {
	result := r.dbClient.Create(&product)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to create product")
	}

	return &product, nil
}
