package repository

import (
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// product repository
type Repository struct {
	dbClient *gorm.DB
}

var (
	// error when db client is nil
	ErrNilDBClient = errors.New("DB client cannot be nil")
)

// creates a new product repository
func NewRepository(dbClient *gorm.DB) (*Repository, error) {
	if dbClient == nil {
		return nil, ErrNilDBClient
	}

	return &Repository{dbClient}, nil
}

// creates a product in the database
func (r *Repository) Create(product entity.Product) (*entity.Product, error) {
	result := r.dbClient.Create(&product)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to create product")
	}

	return &product, nil
}

// gets all products from the database
func (r *Repository) GetAll() ([]*entity.Product, error) {
	products := []*entity.Product{}
	result := r.dbClient.Find(&products)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get all products")
	}

	return products, nil
}

// get a product by id from the database
func (r *Repository) GetByID(id uint) (*entity.Product, error) {
	product := &entity.Product{}
	result := r.dbClient.First(&product, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get product by id")
	}

	return product, nil
}

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

// update a product by id from the database
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
