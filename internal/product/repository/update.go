package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// updates a product by id from the database
func (r *Repository) UpdateByID(id uint, product entity.Product) (*entity.Product, error) {
	p := &entity.Product{}
	result := r.dbClient.First(p, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	result = r.dbClient.Model(p).Updates(product)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to update product by id")
	}

	return p, nil
}
