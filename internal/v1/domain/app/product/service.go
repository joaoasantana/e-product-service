package product

import (
	"errors"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/core/category"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/core/product"
)

const (
	errCategoryNotFound = "category not found"

	errProductAlreadyExists = "product already exists"
	errProductNotFound      = "product not found"
	errProductsNotFound     = "products not found"

	errProductCreationFailed = "product creation failed"
)

type Service struct {
	categoryRepository category.Repository
	productRepository  product.Repository
}

func NewService(categoryRepository category.Repository, productRepository product.Repository) *Service {
	return &Service{
		categoryRepository: categoryRepository,
		productRepository:  productRepository,
	}
}

func (s *Service) Create(input *Input) (string, error) {
	if _, err := s.categoryRepository.FindByID(input.CategoryID); err != nil {
		return "", errors.New(errCategoryNotFound)
	}

	if _, err := s.productRepository.FindByName(input.Name); err == nil {
		return "", errors.New(errProductAlreadyExists)
	}

	entity := &product.Entity{
		Name:        input.Name,
		Description: input.Description,
		CategoryID:  input.CategoryID,
	}

	id, err := s.productRepository.Create(entity)
	if err != nil {
		return "", errors.New(errProductCreationFailed)
	}

	return id, nil
}

func (s *Service) FindAll() ([]Output, error) {
	entities, err := s.productRepository.FindAll()
	if err != nil {
		return nil, errors.New(errProductsNotFound)
	}

	if len(entities) == 0 {
		return nil, errors.New(errProductsNotFound)
	}

	var outputs []Output

	for _, entity := range entities {
		outputs = append(outputs, Output{
			ID:          entity.ID,
			Name:        entity.Name,
			Description: entity.Description,
			CategoryID:  entity.CategoryID,
		})
	}

	return outputs, nil
}

func (s *Service) FindByID(id string) (*Output, error) {
	entity, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, errors.New(errProductNotFound)
	}

	if entity == nil {
		return nil, errors.New(errProductNotFound)
	}

	output := Output{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CategoryID:  entity.CategoryID,
	}

	return &output, nil
}
