package entity

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct_ValidateRules(t *testing.T) {
	t.Run("valid_product", func(t *testing.T) {
		entity := Product{
			Name:       "Unbranded Concrete Cheese",
			CategoryID: "67fac9a48fc1d409fddf131c",
		}

		err := entity.ValidateRules()

		assert.Nil(t, err)
	})

	t.Run("invalid_product_name", func(t *testing.T) {
		entity := Product{
			Name: "",
		}

		err := entity.ValidateRules()

		assert.Equal(t, err, errors.New("name is required"))
	})

	t.Run("invalid_product_category_id", func(t *testing.T) {
		entity := Product{
			Name:       "Unbranded Concrete Cheese",
			CategoryID: "",
		}

		err := entity.ValidateRules()

		assert.Equal(t, err, errors.New("category id is required"))
	})
}
