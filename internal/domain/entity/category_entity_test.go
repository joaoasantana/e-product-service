package entity

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategory_ValidateRules(t *testing.T) {
	t.Run("valid_category", func(t *testing.T) {
		entity := Category{
			Name: "shoes",
		}

		err := entity.ValidateRules()

		assert.Nil(t, err)
	})

	t.Run("invalid_category", func(t *testing.T) {
		entity := Category{
			Name: "",
		}

		err := entity.ValidateRules()

		assert.Equal(t, err, errors.New("name is required"))
	})
}
