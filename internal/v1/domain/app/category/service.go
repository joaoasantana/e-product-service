package category

import (
	"errors"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/core/category"
)

const (
	errCategoryAlreadyExists = "category already exists"
	errCategoryNotFound      = "category not found"
	errCategoriesNotFound    = "categories not found"

	errCategoryCreationFailed = "category creation failed"
)

type Service struct {
	categoryRepository category.Repository
}

func NewService(categoryRepository category.Repository) *Service {
	return &Service{
		categoryRepository: categoryRepository,
	}
}

func (s *Service) Create(input *Input) (string, error) {
	if _, err := s.categoryRepository.FindByName(input.Name); err == nil {
		return "", errors.New(errCategoryAlreadyExists)
	}

	entity := &category.Entity{
		Name: input.Name,
	}

	id, err := s.categoryRepository.Create(entity)
	if err != nil {
		return "", errors.New(errCategoryCreationFailed)
	}

	return id, nil
}

func (s *Service) FindAll() ([]Output, error) {
	entities, err := s.categoryRepository.FindAll()
	if err != nil {
		return nil, errors.New(errCategoriesNotFound)
	}

	if len(entities) == 0 {
		return nil, errors.New(errCategoriesNotFound)
	}

	var outputs []Output

	for _, entity := range entities {
		outputs = append(outputs, Output{
			ID:   entity.ID,
			Name: entity.Name,
		})
	}

	return outputs, nil
}

func (s *Service) FindByID(id string) (*Output, error) {
	entity, err := s.categoryRepository.FindByID(id)
	if err != nil {
		return nil, errors.New(errCategoryNotFound)
	}

	if entity == nil {
		return nil, errors.New(errCategoryNotFound)
	}

	output := Output{
		ID:   entity.ID,
		Name: entity.Name,
	}

	return &output, nil
}
