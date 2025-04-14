package entity

import "errors"

type Category struct {
	ID   string
	Name string
}

func (entity *Category) ValidateRules() error {
	if entity.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
