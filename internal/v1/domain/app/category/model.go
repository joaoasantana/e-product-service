package category

type Input struct {
	Name string `json:"name" binding:"required"`
}

type Output struct {
	ID   string `json:"id"   binding:"required"`
	Name string `json:"name" binding:"required"`
}
