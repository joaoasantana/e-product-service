package model

type (
	ProductInput struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		CategoryID  string `json:"category_id"`
	}

	ProductOutput struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		CategoryID  string `json:"category_id"`
	}
)
