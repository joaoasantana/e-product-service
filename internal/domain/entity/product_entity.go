package entity

import "errors"

type Product struct {
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func (entity *Product) ValidateRules() error {
	if entity.Name == "" {
		return errors.New("name is required")
	}
	if entity.CategoryID == "" {
		return errors.New("category id is required")
	}

	return nil
}
