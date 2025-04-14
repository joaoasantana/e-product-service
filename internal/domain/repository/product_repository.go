package repository

import "github.com/joaoasantana/e-product-service/internal/domain/entity"

type ProductRepository interface {
	Create(*entity.Product) error

	FindAll() ([]entity.Product, error)
	FindById(string) (*entity.Product, error)

	Validate(string) error
}
