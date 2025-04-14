package service

import (
	"errors"
	"github.com/joaoasantana/e-product-service/internal/application/model"
	"github.com/joaoasantana/e-product-service/internal/domain/entity"
	"github.com/joaoasantana/e-product-service/internal/domain/repository"
)

const (
	errCategoryAlreadyExists = "category already exists"
	errCategoryListIsEmpty   = "category list is empty"
	errCategoryNotFound      = "category not found"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(
	categoryRepository repository.CategoryRepository,
) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(input *model.CategoryInput) error {
	if err := s.categoryRepository.Validate(input.Name); err == nil {
		return errors.New(errCategoryAlreadyExists)
	}

	category := &entity.Category{
		Name: input.Name,
	}

	if err := category.ValidateRules(); err != nil {
		return err
	}

	if err := s.categoryRepository.Create(category); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) FindAll() ([]model.CategoryOutput, error) {
	categories, err := s.categoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, errors.New(errCategoryListIsEmpty)
	}

	var result []model.CategoryOutput
	for _, category := range categories {
		result = append(result, model.CategoryOutput{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return result, nil
}

func (s *CategoryService) FindByID(id string) (*model.CategoryOutput, error) {
	category, err := s.categoryRepository.FindById(id)
	if err != nil {
		return nil, errors.New(errCategoryNotFound)
	}

	result := &model.CategoryOutput{
		ID:   category.ID,
		Name: category.Name,
	}

	return result, nil
}
