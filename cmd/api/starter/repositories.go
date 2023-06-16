package starter

import (
	productRepository "github.com/lucasvmiguel/stock-api/internal/product/repository"
)

// repositories is a struct that holds all repositories
type repositories struct {
	product *productRepository.Repository
}

// createRepositories creates all repositories
func (s *Starter) createRepositories() (repositories, error) {
	productRepo, err := productRepository.NewRepository(s.gormDB)
	if err != nil {
		return repositories{}, err
	}

	return repositories{
		product: productRepo,
	}, nil
}
