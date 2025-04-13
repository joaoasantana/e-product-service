package product

type Input struct {
	Name        string `json:"name"         binding:"required"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"  binding:"required"`
}

type Output struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CategoryID  string `json:"category_id"`
}
