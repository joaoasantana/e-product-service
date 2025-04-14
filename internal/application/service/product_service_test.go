package service

import (
	"errors"
	"github.com/joaoasantana/e-product-service/internal/application/model"
	"github.com/joaoasantana/e-product-service/internal/domain/entity"
	test "github.com/joaoasantana/e-product-service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProductService_Create(t *testing.T) {
	categoryRepository := test.NewMockCategoryRepository(t)
	productRepository := test.NewMockProductRepository(t)

	productService := NewProductService(categoryRepository, productRepository)

	t.Run("product_already_exists", func(t *testing.T) {
		productRepository.EXPECT().Validate(mock.Anything).Return(nil).Once()

		err := productService.Create(&model.ProductInput{
			Name:       "",
			CategoryID: "",
		})

		assert.Equal(t, err, errors.New("product already exists"))
	})

	t.Run("product_category_not_found", func(t *testing.T) {
		productRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()
		categoryRepository.EXPECT().FindById(mock.Anything).Return(nil, assert.AnError).Once()

		err := productService.Create(&model.ProductInput{
			Name:       "",
			CategoryID: "",
		})

		assert.Equal(t, err, errors.New("category not found"))
	})

	t.Run("product_invalid_name", func(t *testing.T) {
		productRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()
		categoryRepository.EXPECT().FindById(mock.Anything).Return(nil, nil).Once()

		err := productService.Create(&model.ProductInput{
			Name:       "",
			CategoryID: "",
		})

		assert.Equal(t, err, errors.New("name is required"))
	})

	t.Run("product_invalid_category_id", func(t *testing.T) {
		productRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()
		categoryRepository.EXPECT().FindById(mock.Anything).Return(nil, nil).Once()

		err := productService.Create(&model.ProductInput{
			Name:       "Ergonomic Metal Keyboard",
			CategoryID: "",
		})

		assert.Equal(t, err, errors.New("category id is required"))
	})

	t.Run("product_create_failed", func(t *testing.T) {
		productRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()
		categoryRepository.EXPECT().FindById(mock.Anything).Return(nil, nil).Once()
		productRepository.EXPECT().Create(mock.Anything).Return(assert.AnError).Once()

		err := productService.Create(&model.ProductInput{
			Name:       "shoes",
			CategoryID: "67fac9a48fc1d409fddf131c",
		})

		assert.Equal(t, err, assert.AnError)
	})

	t.Run("product_create_success", func(t *testing.T) {
		productRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()
		categoryRepository.EXPECT().FindById(mock.Anything).Return(nil, nil).Once()
		productRepository.EXPECT().Create(mock.Anything).Return(nil).Once()

		err := productService.Create(&model.ProductInput{
			Name:       "Ergonomic Metal Keyboard",
			CategoryID: "67fac9a48fc1d409fddf131c",
		})

		assert.Nil(t, err)
	})
}

func TestProductService_FindAll(t *testing.T) {
	categoryRepository := test.NewMockCategoryRepository(t)
	productRepository := test.NewMockProductRepository(t)

	productService := NewProductService(categoryRepository, productRepository)

	t.Run("product_find_all_failed", func(t *testing.T) {
		productRepository.EXPECT().FindAll().Return(nil, assert.AnError).Once()

		_, err := productService.FindAll()

		assert.Equal(t, err, assert.AnError)
	})

	t.Run("product_list_empty", func(t *testing.T) {
		var products = make([]entity.Product, 0)

		productRepository.EXPECT().FindAll().Return(products, nil).Once()

		result, err := productService.FindAll()

		assert.Equal(t, len(result), 0)
		assert.Equal(t, err, errors.New("product list is empty"))
	})

	t.Run("product_find_all_success", func(t *testing.T) {
		var products = make([]entity.Product, 3)

		productRepository.EXPECT().FindAll().Return(products, nil).Once()

		result, err := productService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, len(result), 3)
	})
}

func TestProductService_FindByID(t *testing.T) {
	categoryRepository := test.NewMockCategoryRepository(t)
	productRepository := test.NewMockProductRepository(t)

	productService := NewProductService(categoryRepository, productRepository)

	t.Run("product_find_by_id_failed", func(t *testing.T) {
		productRepository.EXPECT().FindById(mock.Anything).Return(nil, assert.AnError).Once()

		_, err := productService.FindByID("")

		assert.Equal(t, err, errors.New("product not found"))
	})

	t.Run("product_find_by_id_success", func(t *testing.T) {
		var product = &entity.Product{}

		productRepository.EXPECT().FindById(mock.Anything).Return(product, nil).Once()

		result, err := productService.FindByID("")

		assert.Nil(t, err)
		assert.Equal(t, result.ID, product.ID)
		assert.Equal(t, result.Name, product.Name)
		assert.Equal(t, result.Description, product.Description)
		assert.Equal(t, result.CategoryID, product.CategoryID)
	})
}
