package service

import (
	"errors"
	"github.com/joaoasantana/e-product-service/internal/application/model"
	"github.com/joaoasantana/e-product-service/internal/domain/entity"
	"github.com/joaoasantana/e-product-service/internal/domain/repository"
)

const (
	errProductAlreadyExists = "product already exists"
	errProductListIsEmpty   = "product list is empty"
	errProductNotFound      = "product not found"
)

type ProductService struct {
	categoryRepository repository.CategoryRepository
	productRepository  repository.ProductRepository
}

func NewProductService(
	categoryRepository repository.CategoryRepository,
	productRepository repository.ProductRepository,
) *ProductService {
	return &ProductService{
		categoryRepository: categoryRepository,
		productRepository:  productRepository,
	}
}

func (s *ProductService) Create(input *model.ProductInput) error {
	if err := s.productRepository.Validate(input.Name); err == nil {
		return errors.New(errProductAlreadyExists)
	}

	if _, err := s.categoryRepository.FindById(input.CategoryID); err != nil {
		return errors.New(errCategoryNotFound)
	}

	product := &entity.Product{
		Name:        input.Name,
		Description: input.Description,
		CategoryID:  input.CategoryID,
	}

	if err := product.ValidateRules(); err != nil {
		return err
	}

	if err := s.productRepository.Create(product); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) FindAll() ([]model.ProductOutput, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New(errProductListIsEmpty)
	}

	var result []model.ProductOutput
	for _, product := range products {
		result = append(result, model.ProductOutput{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			CategoryID:  product.CategoryID,
		})
	}

	return result, nil
}

func (s *ProductService) FindByID(id string) (*model.ProductOutput, error) {
	product, err := s.productRepository.FindById(id)
	if err != nil {
		return nil, errors.New(errProductNotFound)
	}

	result := &model.ProductOutput{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  product.CategoryID,
	}

	return result, nil
}
