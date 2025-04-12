package category

type Repository interface {
	Create(*Entity) (string, error)

	FindAll() ([]Entity, error)

	FindByID(string) (*Entity, error)
	FindByName(string) (*Entity, error)
}
