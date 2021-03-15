package productquery

import "github.com/alejogs4/hn-website/src/products/domain/product"

// ProductQueriesUseCases struct to handle all product get operations
type ProductQueriesUseCases struct {
	productQueries product.QueriesRepository
}

// NewProductQueriesUseCases returns a new instance of ProductQueriesUseCases
func NewProductQueriesUseCases(productQueries product.QueriesRepository) ProductQueriesUseCases {
	return ProductQueriesUseCases{productQueries: productQueries}
}

// GetAllProducts .
func (p ProductQueriesUseCases) GetAllProducts() ([]product.Product, error) {
	return p.productQueries.GetAll()
}

// GetProductByID .
func (p ProductQueriesUseCases) GetProductByID(id string) (product.Product, error) {
	return p.productQueries.GetByID(id)
}
