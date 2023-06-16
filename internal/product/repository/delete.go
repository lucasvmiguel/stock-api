package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// delete a product by id from the database
func (r *Repository) DeleteByID(id uint) (*entity.Product, error) {
	product := &entity.Product{}
	result := r.dbClient.First(product, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	result = r.dbClient.Delete(product, id)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to delete product by id")
	}

	return product, nil
}
