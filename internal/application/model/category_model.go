package model

type (
	CategoryInput struct {
		Name string `json:"name"`
	}

	CategoryOutput struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
