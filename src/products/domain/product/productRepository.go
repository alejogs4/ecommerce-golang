package product

// QueriesRepository abstract methods to define how products will be fetched
type QueriesRepository interface {
	GetAll() ([]Product, error)
	GetByID(id string) (Product, error)
}

// CommandsRepository abstract methods to define how products are gonna be handled
type CommandsRepository interface {
	CreateProduct(product Product) error
	UpdateProduct(product Product) error
	BuyProduct(product Product, quantity int) error
}
