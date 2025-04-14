package service

import (
	"errors"
	"github.com/joaoasantana/e-product-service/internal/application/model"
	"github.com/joaoasantana/e-product-service/internal/domain/entity"
	"github.com/joaoasantana/e-product-service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCategoryService_Create(t *testing.T) {
	categoryRepository := test.NewMockCategoryRepository(t)
	categoryService := NewCategoryService(categoryRepository)

	t.Run("category_already_exists", func(t *testing.T) {
		categoryRepository.EXPECT().Validate(mock.Anything).Return(nil).Once()

		err := categoryService.Create(&model.CategoryInput{
			Name: "shoes",
		})

		assert.Equal(t, err, errors.New("category already exists"))
	})

	t.Run("category_invalid_name", func(t *testing.T) {
		categoryRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()

		err := categoryService.Create(&model.CategoryInput{
			Name: "",
		})

		assert.Equal(t, err, errors.New("name is required"))
	})

	t.Run("category_create_failed", func(t *testing.T) {
		categoryRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()
		categoryRepository.EXPECT().Create(mock.Anything).Return(assert.AnError).Once()

		err := categoryService.Create(&model.CategoryInput{
			Name: "shoes",
		})

		assert.Equal(t, err, assert.AnError)
	})

	t.Run("category_create_success", func(t *testing.T) {
		categoryRepository.EXPECT().Validate(mock.Anything).Return(assert.AnError).Once()
		categoryRepository.EXPECT().Create(mock.Anything).Return(nil).Once()

		err := categoryService.Create(&model.CategoryInput{
			Name: "shoes",
		})

		assert.Nil(t, err)
	})
}

func TestCategoryService_FindAll(t *testing.T) {
	categoryRepository := test.NewMockCategoryRepository(t)
	categoryService := NewCategoryService(categoryRepository)

	t.Run("category_find_all_failed", func(t *testing.T) {
		categoryRepository.EXPECT().FindAll().Return(nil, assert.AnError).Once()

		_, err := categoryService.FindAll()

		assert.Equal(t, err, assert.AnError)
	})

	t.Run("category_list_empty", func(t *testing.T) {
		var categories = make([]entity.Category, 0)

		categoryRepository.EXPECT().FindAll().Return(categories, nil).Once()

		result, err := categoryService.FindAll()

		assert.Equal(t, len(result), 0)
		assert.Equal(t, err, errors.New("category list is empty"))
	})

	t.Run("category_find_all_success", func(t *testing.T) {
		var categories = make([]entity.Category, 3)

		categoryRepository.EXPECT().FindAll().Return(categories, nil).Once()

		result, err := categoryService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, len(result), 3)
	})
}

func TestCategoryService_FindByID(t *testing.T) {
	categoryRepository := test.NewMockCategoryRepository(t)
	categoryService := NewCategoryService(categoryRepository)

	t.Run("category_find_by_id_failed", func(t *testing.T) {
		categoryRepository.EXPECT().FindById(mock.Anything).Return(nil, assert.AnError).Once()

		_, err := categoryService.FindByID("")

		assert.Equal(t, err, errors.New("category not found"))
	})

	t.Run("category_find_by_id_success", func(t *testing.T) {
		var category = &entity.Category{}

		categoryRepository.EXPECT().FindById(mock.Anything).Return(category, nil).Once()

		result, err := categoryService.FindByID("")

		assert.Nil(t, err)
		assert.Equal(t, result.ID, category.ID)
		assert.Equal(t, result.Name, category.Name)
	})
}
