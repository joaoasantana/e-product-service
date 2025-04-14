package repository

import "github.com/joaoasantana/e-product-service/internal/domain/entity"

type CategoryRepository interface {
	Create(*entity.Category) error

	FindAll() ([]entity.Category, error)
	FindById(string) (*entity.Category, error)

	Validate(string) error
}
